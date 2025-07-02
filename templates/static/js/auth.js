// ✅ auth.js (simpan token di cookie)
document.addEventListener("DOMContentLoaded", function () {
  const form = document.getElementById("loginForm");
  form.addEventListener("submit", async function (e) {
    e.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    try {
      const response = await fetch("/api/v1/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password })
      });

      const data = await response.json();

      if (response.ok && data.data && data.data.token) {
        document.cookie = `token=${data.data.token}; path=/`;
        window.location.href = "/dashboard";
      } else {
        document.getElementById("loginError").textContent = data.message || "Login gagal";
        document.getElementById("loginError").classList.remove("d-none");
      }
    } catch (err) {
      alert("Terjadi kesalahan: " + err.message);
    }
  });
});