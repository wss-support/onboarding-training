## Prerequisites
* Ensure you have one of the following Java versions on the computer you are going to run the UA: 
  * Java JDK 8
  * Java JRE 8
  * Java JDK 11
* Additionally you need to have Maven Package Manager installed 

## What to do?
* Download pom.xml to the box you are going to run the UA on
* Run the UA scan with `maven.runPreStep=true` and `-logLevel debug`

## Review
* How many unique Maven dependencies were resolved? [5] 
* How many dependencies are there in the pom.xml? [3] 
* How many libraries do you see in WS UI? [6] Does `junit-platform-engine` dependency appear in the Mend UI results? Why?
* Pull the update request for this scan. Does the update request contain `junit-platform-engine`? Why?
* Rerun the scan in such a way so all the dependencies that are in the pom.xml appear in the Mend UI. What changes did you need to make to get the required result?
* Pull the update request for the latest scan. Does the update request contain `junit-platform-engine`? Why?
