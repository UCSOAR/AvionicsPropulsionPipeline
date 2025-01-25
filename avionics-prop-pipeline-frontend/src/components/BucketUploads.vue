<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { BucketObject } from '../models/BucketObject';
import { endpointMapping } from '../utils/constants';

const objects = reactive<BucketObject[]>([]);
const error = ref<string | null>(null);

const fetchObjects = async () => {
  try {
    const result = await fetch(endpointMapping.getBucketUploadsUrl);

    if (result.ok) {
      const data = await result.json();
      objects.splice(0, objects.length, ...data);
      error.value = null;
    } else {
      error.value = 'Failed to fetch uploaded files.';
    }
  } catch (err) {
    error.value = err as string;
  }
};

const deleteObject = async (filename: string) => {
  // Placeholder logic for deleting a file
  const index = objects.findIndex((file) => file.name === filename);
  if (index !== -1) objects.splice(index, 1);
};

onMounted(() => fetchObjects());
</script>

<template>
  <div class="uploaded-files">
    <h3>Uploaded Files</h3>
    <div v-if="objects.length > 0" class="file-list">
      <div v-for="obj in objects" :key="obj.name" class="file-item">
        <div class="file-info">
          <div class="file-icon"></div>
          <div>
            <p class="file-name">{{ obj.name }}</p>
            <p class="file-size">{{ (obj.size / 1024).toFixed(2) }} KB</p>
          </div>
        </div>
        <button @click="deleteObject(obj.name)" class="delete-button">X</button>
      </div>
    </div>
    <p v-else>No uploaded files yet.</p>
    <p v-if="error" class="error-message">{{ error }}</p>
  </div>
</template>

<style scoped>
.uploaded-files {
  flex: 1;
  padding-left: 20px;
}

.file-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.file-item {
  display: flex;
  align-items: center;
  padding: 10px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.file-info {
  display: flex;
  align-items: center;
  flex: 1;
}

.file-icon {
  width: 40px;
  height: 40px;
  background: #007bff;
  border-radius: 50%;
  margin-right: 12px;
}

.file-name {
  font-weight: 600;
}

.file-size {
  font-size: 0.9em;
  color: #666;
}

.delete-button {
  background: none;
  border: none;
  color: #ff4d4f;
  font-size: 1.2em;
  cursor: pointer;
}
</style>
