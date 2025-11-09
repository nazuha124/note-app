"use client";
import { useState } from "react";
import api from "@/lib/api";
import { useRouter } from "next/navigation";

export default function RegisterPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [msg, setMsg] = useState("");
  const router = useRouter();

  const handleRegister = async () => {
    try {
      await api.post("/register", { email, password });
      setMsg("✅ Register success! Redirecting...");
      setTimeout(() => router.push("/login"), 1000);
    } catch (err: any) {
      setMsg(err?.response?.data?.error || "❌ Failed to register");
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <h1 className="text-xl font-semibold mb-4">Register</h1>
      <input
        type="email"
        placeholder="Email"
        className="border p-2 mb-2 w-64"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <input
        type="password"
        placeholder="Password"
        className="border p-2 mb-2 w-64"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <button onClick={handleRegister} className="bg-blue-600 text-white px-4 py-2 rounded">
        Register
      </button>
      <p className="mt-3 text-gray-600">{msg}</p>
    </div>
  );
}
