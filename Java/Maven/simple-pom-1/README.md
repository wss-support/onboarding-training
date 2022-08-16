## Prerequisites
* Ensure you have one of the following Java versions on the computer you are going to run the UA: 
  * Java JDK 8
  * Java JRE 8
  * Java JDK 11
* Additionally you need to have Maven Package Manager installed 

## What to do?
* Download pom.xml to the box you are going to run the UA on
* Run the UA scan with `-logLevel debug` the following way:
  * Run the scan with OOTB config and note the result? Are there any dependencies resolved by Package Manager Resolution? Why? If yes, did you run any maven commands manually?
  * Run the scan with maven.runPreStep=true and note the results. How long did it take to run the scan this time? Why? 

## Review
* Review the pom.xml. How many dependencies are there?
* Review the project created by the UA in Mend UI. How many dependencies are there? Why is there an extra dependency in Mend UI?
* Why this extra dependency is Unknown and with License Requires Review? 
* What happened in the folder you ran the scan on? What was created there? Why?
* What does target folder contain?
* Do you see the following line in your UA log `Downloaded from central: https://repo.maven.apache.org/maven2/org/apache/maven/doxia/doxia-module-fml/1.0/doxia-module-fml-1.0.jar (19 kB at 29 kB/s)` OR `Downloaded from central: https://repo.maven.apache.org/maven2/org/codehaus/plexus/plexus-components/1.1.7/plexus-components-1.1.7.pom (5.0 kB at 46 kB/s)`? If yes, where were they downloaded to? Why? 

