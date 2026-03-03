"use client";

import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { MoreHorizontal, Pin, Pencil, Trash2, Calendar } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import type { Note } from "@/types";
import { useDeleteNote, useUpdateNote } from "@/hooks/use-notes";

const colorMap: Record<string, string> = {
  "": "bg-card",
  default: "bg-card",
  red: "bg-red-50 dark:bg-red-950/30 border-red-200 dark:border-red-800",
  orange:
    "bg-orange-50 dark:bg-orange-950/30 border-orange-200 dark:border-orange-800",
  yellow:
    "bg-yellow-50 dark:bg-yellow-950/30 border-yellow-200 dark:border-yellow-800",
  green:
    "bg-green-50 dark:bg-green-950/30 border-green-200 dark:border-green-800",
  blue: "bg-blue-50 dark:bg-blue-950/30 border-blue-200 dark:border-blue-800",
  purple:
    "bg-purple-50 dark:bg-purple-950/30 border-purple-200 dark:border-purple-800",
  pink: "bg-pink-50 dark:bg-pink-950/30 border-pink-200 dark:border-pink-800",
  gray: "bg-gray-50 dark:bg-gray-950/30 border-gray-200 dark:border-gray-800",
};

interface NoteCardProps {
  note: Note;
  onEdit?: (note: Note) => void;
}

export function NoteCard({ note, onEdit }: NoteCardProps) {
  const deleteNote = useDeleteNote();
  const updateNote = useUpdateNote();

  const bgColor = colorMap[note.color] || colorMap[""];

  const handlePin = () => {
    updateNote.mutate({
      id: note.id,
      data: { is_pinned: !note.is_pinned },
    });
  };

  return (
    <div
      className={cn(
        "group rounded-xl border p-4 transition-shadow hover:shadow-md cursor-pointer",
        bgColor
      )}
      onClick={() => onEdit?.(note)}
    >
      <div className="flex items-start justify-between mb-2">
        <h3 className="font-medium text-sm line-clamp-1 flex-1">
          {note.title}
        </h3>
        <div className="flex items-center gap-1 shrink-0">
          {note.is_pinned && (
            <Pin className="h-3 w-3 text-amber-500 fill-amber-500" />
          )}
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="h-7 w-7 opacity-100 sm:opacity-0 sm:group-hover:opacity-100"
                onClick={(e) => e.stopPropagation()}
              >
                <MoreHorizontal className="h-3.5 w-3.5" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuItem
                onClick={(e) => {
                  e.stopPropagation();
                  handlePin();
                }}
              >
                <Pin className="mr-2 h-4 w-4" />
                {note.is_pinned ? "Bỏ ghim" : "Ghim"}
              </DropdownMenuItem>
              <DropdownMenuItem
                onClick={(e) => {
                  e.stopPropagation();
                  onEdit?.(note);
                }}
              >
                <Pencil className="mr-2 h-4 w-4" />
                Chỉnh sửa
              </DropdownMenuItem>
              <DropdownMenuItem
                onClick={(e) => {
                  e.stopPropagation();
                  deleteNote.mutate(note.id);
                }}
                className="text-red-600"
              >
                <Trash2 className="mr-2 h-4 w-4" />
                Xoá
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      </div>

      {note.content && (
        <p className="text-xs text-muted-foreground line-clamp-4 mb-3 whitespace-pre-wrap">
          {note.content}
        </p>
      )}

      <div className="flex items-center gap-2 flex-wrap">
        <span className="flex items-center gap-0.5 text-[10px] text-muted-foreground">
          <Calendar className="h-3 w-3" />
          {note.date}
        </span>
        {note.tags?.map((tag) => (
          <Badge key={tag} variant="outline" className="text-[10px]">
            {tag}
          </Badge>
        ))}
      </div>
    </div>
  );
}
