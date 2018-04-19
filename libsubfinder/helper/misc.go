//
// misc.go : contains misc helper function
// Written By : @ice3man (Nizamul Rana)
//
// Distributed Under MIT License
// Copyrights (C) 2018 Ice3man
//

package helper

import (
    "crypto/rand"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

// Current result structure
type Result struct {
    Subdomains []string // Subdomains found
    Error      error    // Any error that has occured
}

// Current Bruteforce structure
type BruteforceResult struct {
    Entity string // Current Subdomain we found
    Error  error  // Error
}

// NewUUID generates a random UUID according to RFC 4122
// Taken from : https://play.golang.org/p/4FkNSiUDMg
//
// Used for bruteforcing and detection of Wildcard Subdomains :-)
func NewUUID() (string, error) {
    uuid := make([]byte, 16)
    n, err := io.ReadFull(rand.Reader, uuid)
    if n != len(uuid) || err != nil {
        return "", err
    }
    // variant bits; see section 4.1.1
    uuid[8] = uuid[8]&^0xc0 | 0x80
    // version 4 (pseudo-random); see section 4.1.3
    uuid[6] = uuid[6]&^0xf0 | 0x40
    return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// Reads a config file from disk and returns Configuration structure
func ReadConfigFile() (configuration *Config, err error) {

    var config Config

    // Get current path
    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

    // Read the file by appending config.json to executable directory
    raw, err := ioutil.ReadFile(dir + "/config.json")
    if err != nil {
        return &config, err
    }

    err = json.Unmarshal(raw, &config)
    if err != nil {
        return &config, err
    }

    return &config, nil
}

<<<<<<< 9ae536175028cdedecd50144bfd7999d5e0e09e6
// Returns unique items in a slice
// Adapted from http://www.golangprograms.com/remove-duplicate-values-from-slice.html
func Unique(elements []string) []string {
=======

// Returns unique items in a slice
// Adapted from http://www.golangprograms.com/remove-duplicate-values-from-slice.html
func Unique(elements []string) []string {  
>>>>>>> Updated Commenting Style and some other misc. changes
    // Use map to record duplicates as we find them.
    encountered := map[string]bool{}
    result := []string{}

    for v := range elements {
        if encountered[elements[v]] == true {
            // Do not add duplicate.
        } else {
            // Record this element as an encountered element.
            encountered[elements[v]] = true
            // Append to result slice.
            result = append(result, elements[v])
        }
    }
    // Return the new slice.
    return result
}

// Returns valid subdomains found ending with target domain
func Validate(state *State, strslice []string) (subdomains []string) {
    for _, entry := range strslice {
        if strings.HasSuffix(entry, state.Domain) {
            subdomains = append(subdomains, entry)
        }
    }

    return subdomains
}
