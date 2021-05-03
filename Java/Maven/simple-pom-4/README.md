## Prerequisites
* Ensure you have one of the following Java versions on the computer you are going to run the UA: 
  * Java JDK 8
  * Java JRE 8
  * Java JDK 11
* Additionally you need to have Maven Package Manager installed 

## What to do?
* Download pom.xml to the computer you are going to run the UA
* Run the UA scan with '-logLevel debug' 

## Review
* Review the pom.xml. How many dependencies are there?
* Review the project created by the UA in WS UI. How many dependencies are there?
* Are all the direct dependencies from pom.xml reflected in WS UI?
* Is there anything in the log that would tell you where the missing dependencies are?
* **Fix your configuration AND/OR settings to see all the direct dependencies in WS UI**
