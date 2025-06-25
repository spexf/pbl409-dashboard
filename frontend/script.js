document.addEventListener("DOMContentLoaded", function () {
  const token = localStorage.getItem("token");

  if (window.location.pathname.includes("dashboard.html")) {
    fetch("http://localhost:8080/api/v1/users", {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
      .then(res => {
        if (!res.ok) {
          window.location.href = "login.html";
        }
        return res.json();
      })
      .then(data => {
        document.getElementById("status").textContent = "Logged in as: " + (data.data[0]?.email || "Unknown");
      });
  }

  const loginForm = document.getElementById("loginForm");
  if (loginForm) {
    loginForm.addEventListener("submit", function (e) {
      e.preventDefault();
      const username = document.getElementById("username").value;
      const password = document.getElementById("password").value;

      fetch("http://localhost:8080/api/v1/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password })
      })
        .then(res => res.json())
        .then(data => {
          if (data.data.token) {
            localStorage.setItem("token", data.data.token);
            window.location.href = "dashboard.html";
          } else {
            alert("Login gagal");
          }
        });
    });
  }
});
