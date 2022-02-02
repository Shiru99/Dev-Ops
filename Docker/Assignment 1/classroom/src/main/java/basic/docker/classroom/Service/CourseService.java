package basic.docker.classroom.Service;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import basic.docker.classroom.Model.Course;
import basic.docker.classroom.Repository.CourseRepository;

@Service
public class CourseService {

    @Autowired
    CourseRepository courseRepository;

    public List<Course> getAllCourses() {
        return (List<Course>) courseRepository.findAll();
    }
}
