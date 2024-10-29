package nmap

import (
	"fmt"
	"os"
	"path/filepath"
)

type Outputter struct {
	File     string
	OutDir   string
	Parsed   *NmapRun
	OutFiles map[string]*os.File
}

func NewOutputter(file string, dir string) (Outputter, error) {

	fd, err := os.OpenFile(file, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return Outputter{}, fmt.Errorf("could not open %s: %w", file, err)
	}

	parsed, err := parse(fd)
	if err != nil {
		return Outputter{}, fmt.Errorf("could not parse %s: %w", file, err)
	}

	return Outputter{
		File:     file,
		OutDir:   dir,
		OutFiles: make(map[string]*os.File),
		Parsed:   parsed,
	}, nil

}

func (o Outputter) Run() error {
	for _, host := range o.Parsed.Hosts {
		for _, port := range host.Ports {
			for _, address := range host.Addresses {

				if port.State != "open" {
					continue
				}
				// output to "all" file first
				err := o.output("all", address.Addr, port.PortId)
				if err != nil {
					return err
				}

				// default to using the detected service name
				if port.Service.Name != "" {

					// collate web
					if port.Service.Name == "http" || port.Service.Name == "https" {
						err := o.output("all_web", address.Addr, port.PortId)
						if err != nil {
							return err
						}
					}

					// output to standard service name
					err := o.output(port.Service.Name, address.Addr, port.PortId)
					if err != nil {
						return err
					}
					continue
				}

				// fallback to port maps if no service name present
				err = o.output(tcpServiceName(port.PortId), address.Addr, port.PortId)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (o Outputter) output(service string, ip string, port int) error {
	filename := filepath.Join(o.OutDir, fmt.Sprintf("service_%s.txt", service))
	filename, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	if _, ok := o.OutFiles[service]; !ok {
		newfd, err := os.Create(filename)
		if err != nil {
			return err
		}
		o.OutFiles[service] = newfd
	}

	line := fmt.Sprintf("%s:%d\n", ip, port)
	n, err := o.OutFiles[service].Write([]byte(line))
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("wrote 0 bytes to %s file", service)
	}

	return nil
}
