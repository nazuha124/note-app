"use client";
import { useState } from "react";
import api from "@/lib/api";
import { useRouter } from "next/navigation";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [msg, setMsg] = useState("");
  const router = useRouter();

  const handleLogin = async () => {
    try {
      const res = await api.post("/login", { email, password });
      localStorage.setItem("token", res.data.token);
      setMsg("✅ Login success!");
      setTimeout(() => router.push("/notes"), 1000);
    } catch (err: any) {
      setMsg(err?.response?.data?.error || "❌ Login failed");
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-black text-white">
      <div className="bg-neutral-900 p-10 rounded-2xl shadow-xl w-80 text-center border border-gray-800">
        <h1 className="text-3xl font-bold mb-6 text-white">Login</h1>

        <input
          type="email"
          placeholder="Email"
          className="bg-neutral-800 text-white border border-gray-700 p-2 mb-3 w-full rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />

        <input
          type="password"
          placeholder="Password"
          className="bg-neutral-800 text-white border border-gray-700 p-2 mb-4 w-full rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />

        <button
          onClick={handleLogin}
          className="bg-blue-600 hover:bg-blue-700 transition w-full py-2 rounded font-semibold"
        >
          Login
        </button>

        <p className="mt-4 text-sm text-gray-400">{msg}</p>
      </div>
    </div>
  );
}
