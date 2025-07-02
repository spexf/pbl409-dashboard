// ✅ agents.js (ambil token dari cookie dan gunakan di semua request)
function getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(';').shift();
  return null;
}

const token = getCookie("token");

if (!token) {
  alert("Token tidak ditemukan. Silakan login ulang.");
  window.location.href = "/";
}

const apiUrl = "/api/v1/wazuh/1/agents";

async function loadAgents() {
  const tableBody = document.querySelector("#agentsTable tbody");
  tableBody.innerHTML = "";

  const response = await fetch(apiUrl, {
    headers: { Authorization: `Bearer ${token}` }
  });
  const result = await response.json();

  if (response.ok && result.data) {
    result.data.forEach(agent => {
      tableBody.innerHTML += `
        <tr>
          <td>${agent.name}</td>
          <td>${agent.ip}</td>
          <td>${agent.status}</td>
          <td>
            <a href="/agent/edit?id=${agent.id}" class="btn btn-primary btn-sm"><i class="fas fa-edit"></i></a>
            <a href="/agent/delete?id=${agent.id}" class="btn btn-danger btn-sm"><i class="fas fa-trash"></i></a>
          </td>
        </tr>`;
    });
  } else {
    tableBody.innerHTML = '<tr><td colspan="4" class="text-center">Tidak ada agent</td></tr>';
  }
}

async function addAgent(name, ip) {
  const response = await fetch(apiUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify({ name, ip, groups: "default" })
  });
  return await response.json();
}

async function deleteAgent(id) {
  const response = await fetch(apiUrl, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify({ agents_id: [id] })
  });
  return await response.json();
}
