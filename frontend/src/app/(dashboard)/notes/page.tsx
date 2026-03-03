"use client";

import { useState } from "react";
import { Plus, Search, FileText } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Skeleton } from "@/components/ui/skeleton";
import { NoteCard } from "@/components/notes/note-card";
import { NoteFormDialog } from "@/components/notes/note-form-dialog";
import { useNotes } from "@/hooks/use-notes";
import { AuthGuard } from "@/components/auth-guard";
import type { Note } from "@/types";

export default function NotesPage() {
  return (
    <AuthGuard>
      <NotesContent />
    </AuthGuard>
  );
}

function NotesContent() {
  const [dialogOpen, setDialogOpen] = useState(false);
  const [editNote, setEditNote] = useState<Note | null>(null);
  const [search, setSearch] = useState("");

  const { data, isLoading } = useNotes({
    search: search || undefined,
    limit: 50,
  });

  const notes = data?.data || [];

  // Group: pinned first
  const pinnedNotes = notes.filter((n) => n.is_pinned);
  const otherNotes = notes.filter((n) => !n.is_pinned);

  const handleEdit = (note: Note) => {
    setEditNote(note);
    setDialogOpen(true);
  };

  const handleClose = () => {
    setDialogOpen(false);
    setEditNote(null);
  };

  return (
    <div className="p-4 lg:p-6 space-y-4">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold">Ghi chú</h1>
          <p className="text-sm text-muted-foreground">
            Quản lý ghi chú cá nhân
          </p>
        </div>
        <Button
          onClick={() => setDialogOpen(true)}
          className="bg-gradient-to-r from-amber-500 to-red-500 text-white"
        >
          <Plus className="mr-2 h-4 w-4" />
          Ghi chú mới
        </Button>
      </div>

      {/* Search */}
      <div className="relative max-w-md">
        <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
        <Input
          placeholder="Tìm kiếm ghi chú..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="pl-9"
        />
      </div>

      {/* Loading */}
      {isLoading && (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-3">
          {Array.from({ length: 8 }).map((_, i) => (
            <Skeleton key={i} className="h-40 rounded-xl" />
          ))}
        </div>
      )}

      {/* Empty state */}
      {!isLoading && notes.length === 0 && (
        <div className="flex flex-col items-center justify-center py-12 text-center">
          <FileText className="h-12 w-12 text-muted-foreground/50 mb-3" />
          <h3 className="text-lg font-medium text-muted-foreground">
            Chưa có ghi chú
          </h3>
          <p className="text-sm text-muted-foreground/70">
            Tạo ghi chú đầu tiên bằng nút &quot;Ghi chú mới&quot;
          </p>
        </div>
      )}

      {/* Pinned notes */}
      {pinnedNotes.length > 0 && (
        <div className="space-y-2">
          <h3 className="text-xs font-semibold uppercase text-muted-foreground tracking-wider">
            📌 Đã ghim
          </h3>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-3">
            {pinnedNotes.map((note) => (
              <NoteCard key={note.id} note={note} onEdit={handleEdit} />
            ))}
          </div>
        </div>
      )}

      {/* Other notes */}
      {otherNotes.length > 0 && (
        <div className="space-y-2">
          {pinnedNotes.length > 0 && (
            <h3 className="text-xs font-semibold uppercase text-muted-foreground tracking-wider">
              Khác
            </h3>
          )}
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-3">
            {otherNotes.map((note) => (
              <NoteCard key={note.id} note={note} onEdit={handleEdit} />
            ))}
          </div>
        </div>
      )}

      <NoteFormDialog
        open={dialogOpen}
        onClose={handleClose}
        editNote={editNote}
      />
    </div>
  );
}
