import { useState } from "react";
import { jwtDecode } from "jwt-decode";

interface TokenPayload {
  username: string;
  role: string;
  exp: number;
}

function App() {
  const [token, setToken] = useState<string | null>(null);
  const [message, setMessage] = useState<string | null>(null);

  const login = async (username: string, password: string) => {
    const res = await fetch("http://localhost:8000/api/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username, password }),
    });
    const data = await res.json();
    if (res.ok) {
      setToken(data.token);
      alert("Login successful!");
    } else {
      alert(data.error);
    }
  };

  const fetchProtected = async (role: string) => {
    const res = await fetch(`http://localhost:8000/api/${role}/dashboard`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    const data = await res.json();
    setMessage(data.message || data.error);
  };

  let decoded: TokenPayload | null = null;
  if (token) {
    decoded = jwtDecode<TokenPayload>(token);
  }

  return (
    <div style={{ padding: 20 }}>
      <h2>Multi-role Auth Example</h2>
      {!token && (
        <>
          <button onClick={() => login("admin", "admin123")}>
            Login as Admin
          </button>
          <button onClick={() => login("user", "user123")}>
            Login as User
          </button>
        </>
      )}

      {token && decoded && (
        <>
          <p>
            Logged in as: {decoded.username} ({decoded.role})
          </p>
          {decoded.role === "admin" && (
            <button onClick={() => fetchProtected("admin")}>
              Go to Admin Dashboard
            </button>
          )}
          {decoded.role === "user" && (
            <button onClick={() => fetchProtected("user")}>
              Go to User Profile
            </button>
          )}
          <p>{message}</p>
        </>
      )}
    </div>
  );
}

export default App;