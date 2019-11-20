package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"io/ioutil"
	"os"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileCreate,
		Read:   resourceFileRead,
		Update: resourceFileUpdate,
		Delete: resourceFileDelete,

		Schema: map[string]*schema.Schema{
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFileCreate(d *schema.ResourceData, m interface{}) error {
	filename := d.Get("filename").(string)
	file, e := os.Create(filename)
	if e != nil {
		return e
	}
	defer file.Close()
	_, e = file.WriteString(d.Get("content").(string))
	if e != nil {
		return e
	}
	d.SetId(filename)
	return resourceFileRead(d, m)
}

func resourceFileRead(d *schema.ResourceData, m interface{}) error {
	bytes, e := ioutil.ReadFile(d.Get("filename").(string))

	if e != nil {
		return e
	}
	_ = d.Set("content", string(bytes))
	return nil
}

func resourceFileUpdate(d *schema.ResourceData, m interface{}) error {
	file, e := os.Create(d.Get("filename").(string))
	if e != nil {
		return e
	}
	defer file.Close()
	_, e = file.WriteString(d.Get("content").(string))
	if e != nil {
		return e
	}
	return resourceFileRead(d, m)
}

func resourceFileDelete(d *schema.ResourceData, m interface{}) error {
	e := os.Remove(d.Get("filename").(string))
	if e != nil {
		return e
	}
	return nil
}
