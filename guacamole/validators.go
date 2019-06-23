package guacamole

import (
	"fmt"
	"strconv"
	"strings"
)

var supportedProtocols = [...]string{"vnc", "rdp", "ssh", "telnet"}

// func connectionProtocolValid(i interface{}, k string) (s []string, es []error) {
// 	v, ok := i.(string)
// 	if !ok {
// 		es = append(es, fmt.Errorf("expected type of %s to be string", k))
// 		return
// 	}
// 	gotit := false
// 	for _, elem := range supportedProtocols {
// 		if elem == strings.ToLower(v) {
// 			gotit = true
// 		}
// 	}
// 	if !gotit {
// 		es = append(es, fmt.Errorf("This protocol is not supported %s", k))
// 		return
// 	}
// 	return
// }

func connectionParentValid(i interface{}, k string) (s []string, es []error) {
	v, ok := i.(string)
	if !ok {
		es = append(es, fmt.Errorf("expected type of %s to be string", k))
		return
	}
	_, invalidint := strconv.Atoi(v)
	if invalidint != nil && v != "ROOT" {

		es = append(es, fmt.Errorf("%s has to be a positive number(supplied as a string) or \"ROOT\" ", k))
		return
	}
	return
}
func connectionProtocolValid(i interface{}, k string) (s []string, es []error) {
	validProtocols := []string{"vnc", "rdp", "ssh", "telnet"}
	v, ok := i.(string)
	if !ok {
		es = append(es, fmt.Errorf("expected type of %s to be string", k))
		return
	}
	found := false
	for _, prot := range validProtocols {
		if strings.ToLower(v) == prot {
			found = true
		}
	}
	if !found {
		es = append(es, fmt.Errorf("%s has to be either one of these: %s ", k, validProtocols))
		return
	}
	return
}
