/* ==========================================================================
   AASTU Academic Management Portal - Interactive Logic
   ========================================================================== */

document.addEventListener('DOMContentLoaded', () => {

  // 1. Navigation Tab Switching
  const navBtns = document.querySelectorAll('.nav-btn');
  const tabPanels = document.querySelectorAll('.tab-panel');

  navBtns.forEach(btn => {
    btn.addEventListener('click', () => {
      navBtns.forEach(b => b.classList.remove('active'));
      tabPanels.forEach(p => p.classList.remove('active'));

      btn.classList.add('active');
      const targetTab = btn.getAttribute('data-tab');
      const panel = document.getElementById(`panel-${targetTab}`);
      if (panel) panel.classList.add('active');
    });
  });

  // 2. Interactive ECTS Credit Calculator for Course Registration
  const courseChecks = document.querySelectorAll('.course-check');
  const selectedEctsLabel = document.getElementById('selected-ects');

  function updateEcts() {
    let totalEcts = 0;
    courseChecks.forEach(check => {
      if (check.checked) {
        totalEcts += parseInt(check.getAttribute('data-ects') || 0);
      }
    });
    if (selectedEctsLabel) {
      selectedEctsLabel.textContent = `Selected ECTS: ${totalEcts} / 30 Max`;
    }
  }

  courseChecks.forEach(check => {
    check.addEventListener('change', updateEcts);
  });

  // 3. Confirm Registration Button Action
  const confirmRegBtn = document.getElementById('confirm-reg-btn');
  if (confirmRegBtn) {
    confirmRegBtn.addEventListener('click', () => {
      let selected = [];
      courseChecks.forEach(check => {
        if (check.checked) {
          selected.push(check.getAttribute('data-title'));
        }
      });

      if (selected.length === 0) {
        alert('Please select at least one course to confirm registration.');
      } else {
        alert(`Successfully registered for ${selected.length} course(s)!\nYour schedule has been updated.`);
      }
    });
  }

});
