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
* What command was executed to get the dependency tree? Why do we use `org.apache.maven.plugins:maven-dependency-plugin:2.8:tree` instead of the regular `dependency:tree` command?
* Where/how can you get a dependency path/hierarchy? Let's assume you are looking for a direct dependency of `zookeeper-jute-3.5.7.jar`, how do you find it?  
* Look at the dependency `curator-client-4.3.0.jar` in teh Library View in Mend UI. What is its "Library Found In:"? Are you scanning this location `~/.m2/repository/org/apache/curator/curator-client/4.3.0/curator-client-4.3.0.jar`, why is it shown as the path for this library? 
* Are there any vulnerabilities in this project?
* Remediate dependency `netty-codec-4.1.45.Final.jar`. Top Fix suggests to upgrade it to `netty-codec-4.1.48.Final.jar`. How do you do that?
