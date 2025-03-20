<<<<<<< HEAD
# ynma-go
=======
## [Yonoma](https://yonoma.io/) Go Email Marketing SDK
### Install
```bash
go get github.com/YonomaHQ/yonoma
```
### Setup
First, you need to get an API key:
```go
client := yonoma.NewClient("api-key") 
```
### Usage
## Send your email:
```go
email := yonoma.Email{
    FromEmail:    "updates@yonoma.io",
    ToEmail:      "email@yourdomain.com",
    Subject:      "Welcome to Yonoma - You're In!",
    MailTemplate: "We're excited to welcome you to Yonoma! Your successful signup marks the beginning of what we hope will be an exceptional journey."
}
response, _ := client.Send(email)
```
## Contacts
#### Create new contact
```go
contactData := yonoma.Contact{
    Email:  "glennjacob@example.com",
    Status: "Subscribed" | "Unsubscribed",
    Data: yonoma.ContactData{
        FirstName: string,
        LastName:  string,
        Phone:     string,
        Gender:    string,
        Address:   string,
        City:      string,
        State:     string,
        Country:   string,
        Zipcode:   string,
    },
}
response, _ := client.CreateContact("List Id", contactData) 
```
#### Update contact
```go
contactData := yonoma.Status{
	Status: "Subscribed" | "Unsubscribed",
}
response, _ := client.UnsubscribeContact("List Id", "Contact Id", contactData)
```
#### Add tag to contact
```go
contactData := yonoma.TagId{
	TagId: "Tag Id",
}
response, _ := client.AddContactTag("Contact Id", contactData)
```
#### Remove tag from contact
```go
contactData := yonoma.TagId{
	TagId: "Tag Id",
}
response, _ := client.RemoveContactTag("Contact Id", contactData)

```
## Managing Tags
#### Create a Tag
```go
tagData := yonoma.Tag{
	Name: "Tag Name",
}
response, _ := client.CreateTag(tagData)
```
#### Update a Tag
```go
tagData := yonoma.Tag{
	Name: "Tag Name",
}
response, _ := client.UpdateTag(tagData)
```
#### Delete a Tag
```go
response, _ := client.DeleteTag("Tag Id")
```
#### Retrieve a Specific Tag
```go
response, _ := client.RetrieveTag("Tag Id")
```
#### List All Tags
```go
response, _ := client.ListTags()
```
## Managing Lists
#### Create a List
```go
listData := yonoma.List{
	Name: "List Name",
}
response, _ := client.CreateList(listData)

```
#### Update a List
```go
listData := yonoma.List{
	Name: "List Name",
}
response, _ := client.UpdateList("List Id", listData)
```
#### Delete a List
```go
response, _ := client.DeleteList("List Id")
```
#### Retrieve a Specific List
```go
response, _ := client.RetrieveList("List Id")
```
#### List All Lists
```go
response, _ := client.ListLists()
```


>>>>>>> f000763 (Initial commit)
