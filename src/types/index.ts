export interface Course {
  code: string;
  title: string;
  ects: number;
  instructor: string;
  status: 'Registered' | 'Pending' | 'Completed';
}

export interface GradeRecord {
  course_code: string;
  course_title: string;
  ects: number;
  grade_letter: string;
  grade_points: number;
  semester: string;
}

export interface StudentProfile {
  id: string;
  full_name: string;
  department: string;
  year_standing: number;
  cumulative_gpa: number;
  completed_ects: number;
  active_courses: Course[];
  grades_history: GradeRecord[];
}
