# url-shortener
Url shortener service written in Golang using Redis and docker
Technology used : Golang, Go-fiber, redis db and docker.

Assignment:
1. Able to create a Restful API where the user can give in the URL which they wish to shorten, they will get a shortened version
2. The shortened version can be used with GET request which will redirect the user to the page (I have implemented the routes, but its not working with gofiber in the system)
  BONUS/EXTRA FEATURES:
1. Instead of using file or in memory, I have implemented Redis DB which is an in-memory database, where the information is stored
2. A particular user (IP address) can make only 10 calls to the API, in 30 mins. And everytime the user calls it, the counter decrements by 1 and resets to 10 after 30 minutes
3. I have dockerized it in 3 parts, one container is API level, which containerizes the API and working code.
4. Another Docker is at the db level, (I have used Redis db), which containerizes the db
5. And the whole thing is encapsulated within docker-compose.yaml file
6. User can also provide their custom short url

WORKING:
1. The application runs when the user calls "docker-compose up -d" command at the root of the project, this sets up the project, creates two separate containers
   which contains db and api respectively and starts them, so the main container contains, two sub working containers
2. User can hit POST/ request with localhost:3000/api/v1 and provide url in the body as JSON as : 
  "url":"any_url"
3. In response to it the user will recieve the shortened url, number of times user can hit the particular api(10 being total and 9 after first trigger), time left for it to 
   reset to 10 (30 mins by default)  
4. User can then use GET/ request and input the recieved URL and it should redirect (GET is not working right now, I tried debugging it, but could not find the cause)
5. After the user triggers the POST API, the data gets stored in the redis data base, which can be viewed
6. To view the data in the redis DB: if docker desktop is running, inside the main container, click the console cli option on db container , and run the following commads:
   For all keys : redis-cli KEYS \*
   To view data of particular key: redis-cli GET <key>
   Here, the data is stored as the main URL against the custom url/id 
    
