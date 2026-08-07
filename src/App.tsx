import React, { useState, useEffect } from 'react';
import { StudentProfile, Course } from './types';

export const App: React.FC = () => {
  const [profile, setProfile] = useState<StudentProfile | null>(null);
  const [activeTab, setActiveTab] = useState<'dashboard' | 'registration' | 'grades'>('dashboard');
  const [selectedECTS, setSelectedECTS] = useState<number>(0);

  useEffect(() => {
    // Fetch profile from Go REST API endpoint
    fetch('/api/v1/student/profile')
      .then((res) => res.json())
      .then((data: StudentProfile) => setProfile(data))
      .catch(() => {
        // Fallback state
        setProfile({
          id: 'ETS0001/15',
          full_name: 'Yared Kinetibeb Tesfaye',
          department: 'Computer Engineering',
          year_standing: 5,
          cumulative_gpa: 3.78,
          completed_ects: 142,
          active_courses: [
            { code: 'CoEng-5101', title: 'Embedded Systems Architecture & Design', ects: 6, instructor: 'Dr. Abebe T.', status: 'Registered' },
            { code: 'CoEng-5102', title: 'Advanced Distributed Systems & Web APIs', ects: 5, instructor: 'Eng. Samuel M.', status: 'Registered' },
            { code: 'CoEng-5103', title: 'Agentic AI & Machine Learning Systems', ects: 6, instructor: 'Dr. Helen G.', status: 'Registered' },
          ],
          grades_history: []
        });
      });
  }, []);

  const handleCourseCheck = (ects: number, isChecked: boolean) => {
    if (isChecked) {
      setSelectedECTS((prev) => prev + ects);
    } else {
      setSelectedECTS((prev) => prev - ects);
    }
  };

  return (
    <div className="aastu-app">
      {/* Top Banner */}
      <div className="top-banner">
        <div className="container banner-flex">
          <span><strong>AASTU Notice:</strong> Go Microservice REST API Backend Active v2.0</span>
          <a href="https://www.aastu.edu.et/" target="_blank" rel="noreferrer">AASTU Official Site</a>
        </div>
      </div>

      {/* Header */}
      <header className="portal-header">
        <div className="container header-flex">
          <div className="portal-brand">
            <div className="brand-logo">AASTU</div>
            <div className="brand-text">
              <span className="univ-name">Addis Ababa Science & Technology University</span>
              <span className="portal-title">Go Microservices Academic System</span>
            </div>
          </div>

          <nav className="portal-nav">
            <button className={`nav-btn ${activeTab === 'dashboard' ? 'active' : ''}`} onClick={() => setActiveTab('dashboard')}>
              Dashboard
            </button>
            <button className={`nav-btn ${activeTab === 'registration' ? 'active' : ''}`} onClick={() => setActiveTab('registration')}>
              Course Registration
            </button>
            <button className={`nav-btn ${activeTab === 'grades' ? 'active' : ''}`} onClick={() => setActiveTab('grades')}>
              Grades & Transcript
            </button>
          </nav>
        </div>
      </header>

      {/* Main Container */}
      <main className="container main-content">
        {profile && (
          <div className="student-hero-card">
            <h2>Welcome back, {profile.full_name} 👋</h2>
            <p>ID: {profile.id} | Year: {profile.year_standing}th Year Senior | GPA: {profile.cumulative_gpa} / 4.0 | Completed: {profile.completed_ects} ECTS</p>
          </div>
        )}

        {/* Tab 1: Dashboard */}
        {activeTab === 'dashboard' && profile && (
          <div className="card-panel">
            <h3>Active Enrolled Courses (Go API Connected)</h3>
            <div className="course-list">
              {profile.active_courses.map((course: Course) => (
                <div key={course.code} className="course-item">
                  <h4>{course.code}: {course.title}</h4>
                  <p>Instructor: {course.instructor} | ECTS: {course.ects}</p>
                </div>
              ))}
            </div>
          </div>
        )}

        {/* Tab 2: Course Registration */}
        {activeTab === 'registration' && (
          <div className="card-panel">
            <h3>Course Registration (Selected ECTS: {selectedECTS} / 30 Max)</h3>
            <p>Select senior Computer Engineering courses to calculate registration weight:</p>
            <label>
              <input type="checkbox" onChange={(e) => handleCourseCheck(6, e.target.checked)} /> CoEng-5105 Real-Time Operating Systems (RTOS) (6 ECTS)
            </label>
          </div>
        )}
      </main>
    </div>
  );
};

export default App;
