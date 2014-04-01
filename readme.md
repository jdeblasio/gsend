**Usage**


```
m := &gsend.Gmail {
    From: "from@gmail.com",
    To: "to@somewhere.com",
    Subject: "mail title",
    Message: "body text",
    Password: "<password>",
}

e := m.Send() 
if e != nil {
    panic(e)
}
```
