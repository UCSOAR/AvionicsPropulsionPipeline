<script setup lang="ts">
import { ref, inject } from 'vue';
import { endpointMapping } from '../utils/constants';

const isDarkMode = inject('isDarkMode', ref(false));
const selectedFiles = ref<File[]>([]);
const showTray = ref(false);
const uploadMessage = ref<string>('');  // New reactive variable for the toast message

const onFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  const files = target.files ? Array.from(target.files) : [];
  selectedFiles.value.push(...files);
  showTray.value = true;
};

const onDragOver = (event: DragEvent) => {
  event.preventDefault();
};

const onDrop = (event: DragEvent) => {
  event.preventDefault();
  const files = event.dataTransfer?.files ? Array.from(event.dataTransfer.files) : [];
  selectedFiles.value.push(...files);
  showTray.value = true;
};

const removeFile = (index: number) => {
  selectedFiles.value.splice(index, 1);
  if (selectedFiles.value.length === 0) showTray.value = false;
};

const uploadFiles = async () => {
  if (selectedFiles.value.length === 0) {
    // Optionally, you can also show a toast message for no files selected.
    return;
  }

  for (const file of selectedFiles.value) {
    const formData = new FormData();
    formData.append('file', file);

    try {
      const response = await fetch(endpointMapping.uploadStaticFireUrl, {
        method: 'POST',
        body: formData,
      });

      if (response.ok) {
        console.log(`File ${file.name} uploaded successfully`);
      } else {
        console.error(`Failed to upload file ${file.name}: ${await response.text()}`);
      }
    } catch (error) {
      console.error(`Error uploading file ${file.name}:`, error);
    }
  }

  selectedFiles.value = [];
  showTray.value = false;
  
  // Set the toast message and clear it after 3 seconds
  uploadMessage.value = 'All files uploaded successfully!';
  setTimeout(() => {
    uploadMessage.value = '';
  }, 3000);
};
</script>

<template>
  <div>
    <div 
      :class="['upload-button-wrapper', { dark: isDarkMode }]" 
      @dragover="onDragOver" 
      @drop="onDrop"
    >
      <button type="button" :class="['upload-button', { dark: isDarkMode }]">
        Browse Files
        <input type="file" multiple @change="onFileChange" class="file-input" />
      </button>
    </div>

    <div v-if="showTray" :class="['file-tray', { dark: isDarkMode }]">
      <ul class="file-list">
        <li v-for="(file, index) in selectedFiles" :key="index" class="file-item">
          <div class="file-info">
            <p class="file-name">{{ file.name }}</p>
            <p class="file-size">{{ (file.size / 1024).toFixed(2) }} KB</p>
          </div>
          <button @click="removeFile(index)" class="file-remove">X</button>
        </li>
      </ul>
      <button @click="uploadFiles" :class="['upload-tray-button', { dark: isDarkMode }]">
        Confirm Upload
      </button>
    </div>

    <!-- Toast notification for upload success -->
    <div v-if="uploadMessage" :class="['upload-message', { dark: isDarkMode }]">
      {{ uploadMessage }}
    </div>
  </div>
</template>

<style scoped lang="scss">
@import '../styles/variables.scss';
/* Base Light Theme */
.upload-button-wrapper {
  display: flex;
  justify-content: flex-start;
  padding: 1rem;
}

.upload-button {
  background: $soar-red-color;
  color: #fff;
  padding: 10px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  position: relative;
  width: 240px;
}

.upload-button:hover {
  background: $soar-red-color;
}

.upload-button input {
  position: absolute;
  top: 0;
  left: 0;
  opacity: 0;
  width: 100%;
  height: 100%;
  cursor: pointer;
}

/* Dark Mode */
.upload-button.dark {
  background: #444;
  color: #fff;
}

.upload-button.dark:hover {
  background: #666;
}

/* File Tray */
.file-tray {
  background: #fff;
  border-top: 2px solid #007bff;
  padding: 10px 0px;
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  box-shadow: 0 -4px 6px rgba(0, 0, 0, 0.1);
  transform: translateY(0);
  animation: slide-in 0.4s ease-out;
  z-index: 1000;
}

/* Dark Mode File Tray */
.file-tray.dark {
  background: #333;
  border-color: #555;
  color: #fff;
}

@keyframes slide-in {
  from {
    transform: translateY(100%);
  }
  to {
    transform: translateY(0);
  }
}

.file-list {
  list-style: none;
  padding: 0;
  margin: 0 0 20px;
}

.file-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px;
  margin-bottom: 10px;
  background: #f9f9f9;
  border-radius: 5px;
}

/* Dark Mode File Items */
.file-tray.dark .file-item {
  background: #444;
}

.file-info {
  display: flex;
  flex-direction: column;
}

.file-name {
  font-weight: 600;
}

.file-size {
  font-size: 12px;
  color: #666;
}

/* Dark Mode Text */
.file-tray.dark .file-size {
  color: #ccc;
}

.file-remove {
  background: transparent;
  border: none;
  color: #ff4d4f;
  cursor: pointer;
  font-size: 16px;
}

/* Dark Mode Remove Button */
.file-tray.dark .file-remove {
  color: #ff7777;
}

.upload-tray-button {
  background: #007bff;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-weight: 600;
  display: block;
  margin: 0 auto;
}

/* Dark Mode Upload Button */
.upload-tray-button.dark {
  background: #555;
  color: #fff;
}

.upload-tray-button.dark:hover {
  background: #777;
}

/* Toast Notification (Upload Message) */
.upload-message {
  position: fixed;
  bottom: 90px; /* positioned just above the file tray */
  left: 50%;
  transform: translateX(-50%);
  background: #007bff;
  color: white;
  padding: 10px 20px;
  border-radius: 5px;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
  opacity: 0.95;
  z-index: 1001;
  transition: opacity 0.3s ease;
}

/* Dark Mode Toast */
.upload-message.dark {
  background: #555;
  color: #fff;
}
</style>