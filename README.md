# GSES BTC application

GSES BTC application provides the API for retrieving the current exchange rate of Dollar (USD) in Ukrainian Hryvnia (UAH) and managing subscriptions for rate updates via email

# How to Use

####  For using base URL: gses2.app/api  
- add gses2.app to etc/hosts by command:
####
       sudo sh -c 'echo "127.0.0.1 gses2.app" >> /etc/hosts'

- Create proxy redirect API from 80 to 8000.   Example for nginx:
####    
      sudo sh -c 'echo "server {
      listen 127.0.0.1:80;
      server_name gses2.app;
          location /api/rate {
              proxy_pass http://localhost:8000/api/rate;
              proxy_set_header Host $host;
              proxy_set_header X-Real-IP $remote_addr;
              proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
              proxy_set_header X-Forwarded-Proto $scheme;
          }
    
          location /api/subscribe {
              proxy_pass http://localhost:8000/api/subscribe;
              proxy_set_header Host $host;
              proxy_set_header X-Real-IP $remote_addr;
              proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
              proxy_set_header X-Forwarded-Proto $scheme;
          }
    
          location / {
    
          } 
        } " >> /etc/nginx/sites-available/gses2.app'

####      
        ln -s /etc/nginx/sites-available/gses2.app /etc/nginx/sites-enabled
####
        sudo systemctl restart nginx

#### For configure app replace you smtp host, email name and password  in .env

        HOST_EMAIL = "smtp.host.com"
        SERVICE_EMAIL = your_service_email@example.com
        SERVICE_EMAIL_PASSWD = your_service_email_password

## Launch the app by the following commands:

    make build
    make up
     
## API Endpoints

#### Get Current Exchange Rate
- **URL**: http://gses2.app/api/rate
- **Method:** GET
- **Description:** Retrieve the current exchange rate
- **Request Body:** Returns the current exchange rate .
- **CLI example:** curl http://gses2.app/api/rate

#### Subscribe to the newsletter
- **URL**: http://gses2.app/api/subscribe
- **Method:** POST
- **Description:** Subscribe an email address to receive rate update.
- **Request Body:** Pass the email address in the request body as email.
- **Response:** Returns a success message if the email address was saved.
- **CLI example:** curl -X POST -d "email=test@test.com" http://gses2.app/api/subscribe

