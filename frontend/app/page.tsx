"use client";
import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-black text-white">
      <main className="text-center">
        <h1 className="text-4xl font-bold mb-4">Selamat Datang 👋</h1>
        <p className="text-gray-400 mb-10 max-w-md">
          Aplikasi Catatan — Simpan dan kelola catatanmu dengan mudah 📓
        </p>

        <div className="flex flex-col sm:flex-row gap-4 justify-center">
          <button
            onClick={() => router.push("/login")}
            className="bg-blue-600 hover:bg-blue-700 transition px-6 py-3 rounded-lg font-semibold"
          >
            Login
          </button>

          <button
            onClick={() => router.push("/register")}
            className="bg-gray-700 hover:bg-gray-800 transition px-6 py-3 rounded-lg font-semibold"
          >
            Register
          </button>
        </div>
      </main>
    </div>
  );
}
