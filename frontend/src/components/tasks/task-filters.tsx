"use client";

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Search } from "lucide-react";
import type { TaskStatus, TaskPriority } from "@/types";

interface TaskFiltersProps {
  status: string;
  priority: string;
  search: string;
  onStatusChange: (status: string) => void;
  onPriorityChange: (priority: string) => void;
  onSearchChange: (search: string) => void;
}

export function TaskFilters({
  status,
  priority,
  search,
  onStatusChange,
  onPriorityChange,
  onSearchChange,
}: TaskFiltersProps) {
  return (
    <div className="flex flex-col sm:flex-row gap-2">
      <div className="relative flex-1">
        <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
        <Input
          placeholder="Tìm kiếm..."
          value={search}
          onChange={(e) => onSearchChange(e.target.value)}
          className="pl-9"
        />
      </div>

      <Select value={status} onValueChange={onStatusChange}>
        <SelectTrigger className="w-full sm:w-40">
          <SelectValue placeholder="Trạng thái" />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="all">Tất cả</SelectItem>
          <SelectItem value="todo">Cần làm</SelectItem>
          <SelectItem value="in_progress">Đang làm</SelectItem>
          <SelectItem value="done">Hoàn thành</SelectItem>
          <SelectItem value="cancelled">Đã hủy</SelectItem>
        </SelectContent>
      </Select>

      <Select value={priority} onValueChange={onPriorityChange}>
        <SelectTrigger className="w-full sm:w-40">
          <SelectValue placeholder="Ưu tiên" />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="all">Tất cả</SelectItem>
          <SelectItem value="urgent">🔴 Khẩn cấp</SelectItem>
          <SelectItem value="high">🟠 Cao</SelectItem>
          <SelectItem value="medium">🔵 Trung bình</SelectItem>
          <SelectItem value="low">🟢 Thấp</SelectItem>
        </SelectContent>
      </Select>
    </div>
  );
}
