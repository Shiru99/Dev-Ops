package basic.docker.classroom.Controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import basic.docker.classroom.Model.Course;
import basic.docker.classroom.Service.CourseService;

@RestController
@RequestMapping("/api/classroom")
public class ClassroomController {

    @Value("${spring.application.name}")
    private String appName;

    @Value("${schedule.greetings}")
    private String greetings;

    @GetMapping("/")
    public String index() {
        return greetings + " from " + appName;
    }

    @Autowired
    CourseService courseService;

    @GetMapping(value = "/getCourses", produces = "application/json")
    public List<Course> getAllCourses() {
        return courseService.getAllCourses();
    }
}
