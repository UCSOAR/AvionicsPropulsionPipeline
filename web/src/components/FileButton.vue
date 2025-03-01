<script setup lang="ts">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faTrash, faFile } from "@fortawesome/free-solid-svg-icons";
import { PreviewMetadata } from "../models/PreviewMetadata";

const props = defineProps<{
  metadata: PreviewMetadata;
  fileName: string;
  isDarkMode: boolean;
}>();

const emit = defineEmits<{
  (e: 'fileClick', metadata: PreviewMetadata, fileName: string): void;
  (e: 'deleteClick', fileName: string): void;
}>();

// Helper function to remove the extension from a file name
const removeExtension = (fileName: string): string => {
  const lastDotIndex = fileName.lastIndexOf(".");
  return lastDotIndex !== -1 ? fileName.substring(0, lastDotIndex) : fileName;
};
</script>

<template>
  <div :class="['file-item', { dark: isDarkMode }]">
    <button type="button" class="file-info-button"
      style="background: none; border: none; padding: 0; margin: 0; cursor: pointer;"
      @click="emit('fileClick', metadata, fileName)">
      <div class="file-info">
        <div class="file-icon">
          <font-awesome-icon :icon="faFile" />
        </div>
        <div class="button-name">
          <p class="file-name">{{ removeExtension(fileName) }}</p>
        </div>
      </div>
    </button>
    <button type="button" @click="emit('deleteClick', fileName)" class="delete-button">
      <font-awesome-icon :icon="faTrash" />
    </button>
  </div>
</template>

<style scoped>
/* File Item */
.file-item {
  gap: 10px;
  justify-content: space-between;
  display: flex;
  height: 30px;
  align-items: center;
  padding: 10px;
  background: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  transition: background 0.3s ease;
}

.file-item.dark {
  background: #333;
  color: #fff;
}

/* File Info */
.file-info {
  display: flex;
  align-items: center;
  flex: 1;
}

/* File Icon */
.file-icon {
  color: #007bff;
  border-radius: 50%;
  margin-right: 12px;
}

.file-item.dark .file-icon {
  color: #66b2ff;
}

/* File Name */
.file-name {
  font-weight: 500;
  color: #333;
  margin: 0;
}

.file-item.dark .file-name {
  color: #fff;
}

/* Delete Button */
.delete-button {
  background: none;
  border: none;
  color: #ff4d4f;
  font-size: 1.2em;
  cursor: pointer;
  transition: color 0.2s ease;
}

.delete-button:hover {
  color: #d63031;
}

.file-item.dark .delete-button {
  color: #ff7777;
}

.file-item.dark .delete-button:hover {
  color: #ff9999;
}
</style>