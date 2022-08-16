## Prerequisites
* Ensure you have one of the following Java versions on the computer you are going to run the UA: 
  * Java JDK 8
  * Java JRE 8
  * Java JDK 11
* Additionally you need to have Maven Package Manager installed 

## What to do?
* Download pom.xml to the computer you are going to run the UA
* Run the UA scan with '-logLevel debug' 

## What are the results?
* Review the pom.xml. How many dependencies are there?
* Review the project created by the UA in Mend UI. How many dependencies are there? Why there is a difference in the number of the dependencies in pom.xml and in the number of the dependencies in Mend UI?
* Are there any vulnerabilities?
* Remediate dependency ## netty-codec-4.1.45.Final.jar. Top Fix suggests to upgrade it to ## netty-codec-4.1.48.Final.jar
