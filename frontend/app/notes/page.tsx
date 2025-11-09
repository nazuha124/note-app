"use client";
import { useEffect, useState } from "react";
import api from "@/lib/api";

export default function NotesPage() {
  const [notes, setNotes] = useState<any[]>([]);
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [image, setImage] = useState<File | null>(null);
  const [msg, setMsg] = useState("");

  // ambil data notes dari backend
  const fetchNotes = async () => {
    try {
      const res = await api.get("/api/notes");
      setNotes(res.data);
    } catch (err: any) {
      setMsg(err?.response?.data?.error || "Gagal ambil data");
    }
  };

  useEffect(() => {
    fetchNotes();
  }, []);

  // tambah catatan baru
  const handleAddNote = async () => {
    const formData = new FormData();
    formData.append("title", title);
    formData.append("content", content);
    if (image) formData.append("image", image);

    try {
      await api.post("/api/notes", formData, {
        headers: { "Content-Type": "multipart/form-data" },
      });
      setTitle("");
      setContent("");
      setImage(null);
      fetchNotes();
      setMsg("✅ Catatan berhasil ditambahkan!");
    } catch (err: any) {
      setMsg(err?.response?.data?.error || "Gagal menambah catatan");
    }
  };

  // hapus catatan
  const handleDelete = async (id: number) => {
    try {
      await api.delete(`/api/notes/${id}`);
      fetchNotes();
    } catch {
      alert("Gagal menghapus catatan");
    }
  };

  return (
    <div className="min-h-screen bg-black text-white p-6">
      <h1 className="text-3xl font-bold mb-6 text-center">📓 Catatan Saya</h1>

      <div className="bg-neutral-900 p-6 rounded-2xl shadow-xl max-w-lg mx-auto mb-10 border border-gray-800">
        <h2 className="text-xl mb-3 font-semibold">Tambah Catatan</h2>
        <input
          type="text"
          placeholder="Judul"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          className="bg-neutral-800 text-white p-2 w-full mb-3 rounded border border-gray-700"
        />
        <textarea
          placeholder="Isi catatan"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          className="bg-neutral-800 text-white p-2 w-full mb-3 rounded border border-gray-700"
        />
        <input
          type="file"
          onChange={(e) => setImage(e.target.files?.[0] || null)}
          className="mb-3"
        />
        <button
          onClick={handleAddNote}
          className="bg-blue-600 hover:bg-blue-700 py-2 px-4 rounded font-semibold"
        >
          Simpan
        </button>
        <p className="mt-3 text-gray-400 text-sm">{msg}</p>
      </div>

      <div className="max-w-3xl mx-auto grid gap-4 sm:grid-cols-2">
        {Array.isArray(notes) && notes.length > 0 ? (
            notes.map((note) => (
            <div
                key={note.id}
                className="bg-neutral-900 p-4 rounded-xl border border-gray-800 shadow-md"
            >
                <h3 className="text-lg font-bold mb-1">{note.title}</h3>
                <p className="text-gray-300 mb-2">{note.content}</p>
                {note.image_url && (
                <img
                    src={`http://localhost:8081/${note.image_url}`}
                    alt="note"
                    className="rounded-md mb-2 max-h-64 w-auto object-contain"
                />
                )}
                <button
                onClick={() => handleDelete(note.id)}
                className="bg-red-600 hover:bg-red-700 py-1 px-3 rounded text-sm"
                >
                Hapus
                </button>
            </div>
            ))
        ) : (
            <p className="text-center text-gray-500 col-span-2">
            Tidak ada catatan yang tersedia.
            </p>
        )}
    </div>

    </div>
  );
}

