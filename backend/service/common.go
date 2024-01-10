package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	DB_FILE_PREFIX    = "./jdrnotes-"
	DB_FILE_EXTENSION = ".json"
	USER_TABLE        = "user"
	CAMPAIGN_TABLE    = "campaign"
	CONTAINER_TABLE   = "container"
	ELEMENT_TABLE     = "element"
	TAG_TABLE         = "tag"
	CATEGORY_TABLE    = "category"
	FILE_TABLE        = "file"
	ICON_TABLE        = "icon"
)

var lock sync.Mutex

// Marshal is a function that marshals the object into an
// io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// Save saves a representation of v to the file at path.
func Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func Load(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	reader, err := os.Open(path)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist: ", err)
		return nil
	}
	if err != nil {
		fmt.Println("File exist but: ", err)
		return err
	}
	defer reader.Close()
	if err := Unmarshal(reader, &v); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Load succeed, Object: ", v)
	return nil
}
