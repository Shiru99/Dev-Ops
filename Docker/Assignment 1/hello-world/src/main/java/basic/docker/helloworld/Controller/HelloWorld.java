package basic.docker.helloworld.Controller;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/hello-world")
public class HelloWorld {

    @Value("${spring.application.name}")
    private String appName;

    @Value("${hello.greetings}")
    private String greetings;
    
    @GetMapping("/")
    public String index() {
        return greetings+" from "+appName;
    }
}
