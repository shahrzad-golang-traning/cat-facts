# cat-facts


# 🐱 Cat Fact of the Day - Go Edition  
Ever wondered how quirky cats really are? This simple Go app brings you a new random cat fact every time you run it – straight from the internet! No fluff, just facts.  
This command-line app fetches a random cat fact from the public [catfact.ninja](https://catfact.ninja) API and prints it right to your terminal. Whether you're a cat lover, Go newbie, or just bored – this one's for you.  
Built with Go (Golang), using `net/http` for making HTTP requests and `encoding/json` for decoding API responses.  
To run it, you’ll need Go 1.16 or higher installed on your machine.  

Clone the repository:  
```bash  
git clone https://github.com/yourusername/cat-fact-go.git  
cd cat-fact-go  
```  

Then simply run:  
```bash  
go run main.go  
```  

You’ll see something like:  
```
============= API Response =============
Status: 200 OK
Status Code: 200
Headers: application/json
----------------------------------------
Body:
{"fact":"In the 1750s, Europeans introduced cats into the Americas to control pests.","length":75}
```  

Why?  
Because cats rule the internet. And Go makes things fast and fun.  

API Reference: [catfact.ninja](https://catfact.ninja)  

Future Ideas:  
- Save facts to a local file  
- Set up a daily notification using Task Scheduler or cron  
- Send facts via SMS/email using Twilio or Mailgun  

License: MIT – do whatever you want, just don’t forget to feed your cat.
