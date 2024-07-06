package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "github.com/pquerna/otp/totp"
)

type User struct {
    Username string
    Password string
    Secret   string
}