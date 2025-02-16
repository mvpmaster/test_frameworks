


https://stackoverflow.com/questions/34966648/java-home-is-not-defined-correctly-on-ubuntu
https://maven.apache.org/guides/getting-started/maven-in-five-minutes.html
https://stackoverflow.com/questions/29920434/maven-adding-mainclass-in-pom-xml-with-the-right-folder-path

mvn archetype:generate -DarchetypeGroupId=io.activej -DarchetypeArtifactId=archetype-http -DarchetypeVersion=5.4.3

mvn --version
mvn package

#java -cp target/my-web-app-5.4.3.jar org.example.WebApp
#mvn compile exec:java -Dexec.mainClass="java.org.example.WebApp.main.Exec"

    <configuration>
        <archive>
        <manifest>
            <addClasspath>true</addClasspath>
            <mainClass>org.example.WebApp</mainClass>
        </manifest>
        </archive>
    </configuration>


mvn compile exec:java -Dexec.mainClass="org.example.WebApp"


java -jar ./target/my-web-app-5.4.3.jar