<script setup lang="ts">
import { ref, watch } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faCloudArrowUp } from '@fortawesome/free-solid-svg-icons';

const selectedFiles = ref<File[]>([]);
const showTray = ref(false);

const onFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  const files = target.files ? Array.from(target.files) : [];
  selectedFiles.value.push(...files);
  showTray.value = true; // Show the tray when files are added
};

const onDragOver = (event: DragEvent) => {
  event.preventDefault();
};

const onDrop = (event: DragEvent) => {
  event.preventDefault();
  const files = event.dataTransfer?.files ? Array.from(event.dataTransfer.files) : [];
  selectedFiles.value.push(...files);
  showTray.value = true; // Show the tray when files are dropped
};

const removeFile = (index: number) => {
  selectedFiles.value.splice(index, 1);
  if (selectedFiles.value.length === 0) showTray.value = false; // Hide the tray if no files remain
};

const uploadFiles = async () => {
  if (selectedFiles.value.length === 0) {
    alert('Please select files to upload.');
    return;
  }

  for (const file of selectedFiles.value) {
    const formData = new FormData();
    formData.append('file', file);

    try {
      const response = await fetch('http://localhost:8080/BucketUpload', {
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

  // Clear files after successful upload
  selectedFiles.value = [];
  showTray.value = false;
  alert('All files uploaded successfully!');
};

</script>

<template>
  <div class="upload-form">
    <!-- Drag and Drop Area -->
    <div 
      class="upload-box" 
      @dragover="onDragOver" 
      @drop="onDrop"
    >
      <FontAwesomeIcon :icon="faCloudArrowUp" class="upload-icon" />
      <p>Drag and drop files here</p>
      <button type="button" class="upload-button">
        Browse Files
        <input type="file" multiple @change="onFileChange" class="file-input" />
      </button>
    </div>

    <!-- Sliding Tray for Uploaded Files -->
    <div v-if="showTray" class="file-tray">
      <ul class="file-list">
        <li v-for="(file, index) in selectedFiles" :key="index" class="file-item">
          <div class="file-info">
            <div class="file-icon"></div>
            <div>
              <p class="file-name">{{ file.name }}</p>
              <p class="file-size">{{ (file.size / 1024).toFixed(2) }} KB</p>
            </div>
          </div>
          <button @click="removeFile(index)" class="remove-button">X</button>
        </li>
      </ul>
      <button @click="uploadFiles" class="upload-tray-button">Upload All Files</button>
    </div>
  </div>
</template>

<style scoped>
/* Font Import */
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;600&display=swap');

/* Base Styles */
.upload-form {
  font-family: 'Poppins', sans-serif;
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 600px;
  margin: auto;
}

.upload-box {
  width: 100%;
  height: 200px;
  border: 2px dashed #007bff;
  border-radius: 10px;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
  background: #f9f9f9;
  transition: background-color 0.3s ease;
}

.upload-box:hover {
  background: #e8f0fe;
}

.upload-icon {
  font-size: 40px;
  color: #007bff;
  margin-bottom: 10px;
}

.upload-button {
  background: #007bff;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-weight: 600;
  position: relative;
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

/* Sliding Tray Styles */
.file-tray {
  width: 100%;
  background: #fff;
  border-top: 2px solid #007bff;
  box-shadow: 0 -4px 6px rgba(0, 0, 0, 0.1);
  position: fixed;
  bottom: 0;
  left: 0;
  padding: 20px;
  transform: translateY(0);
  animation: slide-in 0.4s ease-out;
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

.file-info {
  display: flex;
  align-items: center;
}

.file-icon {
  width: 24px;
  height: 24px;
  background-color: #d8d8d8;
  border-radius: 50%;
  margin-right: 10px;
}

.file-name {
  font-weight: 600;
}

.file-size {
  font-size: 12px;
  color: #666;
}

.remove-button {
  background: transparent;
  border: none;
  color: #ff4d4f;
  cursor: pointer;
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
</style>
