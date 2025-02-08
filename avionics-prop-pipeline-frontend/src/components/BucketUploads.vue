<script setup lang="ts">
import { onMounted, reactive, ref, inject } from "vue";
import { PreviewMetadata } from "../models/BucketObject";
import { endpointMapping } from "../utils/constants";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faTrash, faFile } from "@fortawesome/free-solid-svg-icons";
import { useMetadataStore } from "../stores/metadataStore";

const isDarkMode = inject("isDarkMode", ref(false));
const cachePreviews = reactive<Record<string, PreviewMetadata>>({});
const metadataStore = useMetadataStore();
const error = ref<string | null>(null);

const fetchCachePreviews = async () => {
  try {
    const result = await fetch(endpointMapping.getStaticFireMetadataUrl);
    if (result.ok) {
      const data = await result.json();
      console.log(data);
      Object.assign(cachePreviews, data);
      error.value = null;
    } else {
      error.value = "Failed to fetch uploaded files.";
    }
  } catch (err) {
    error.value = err as string;
  }
};

const handleFileClick = (obj: PreviewMetadata) => {
  metadataStore.setMetadata(obj, obj.xColumnNames[0], obj.yColumnNames[0]);
  console.log('meta data stored', obj)
};

const deleteCache = async (cacheName: string) => {
  if (cacheName in cachePreviews) {
    delete cachePreviews[cacheName];
  }
};

// Helper function to remove the extension from a file name
const removeExtension = (fileName: string): string => {
  const lastDotIndex = fileName.lastIndexOf(".");
  return lastDotIndex !== -1 ? fileName.substring(0, lastDotIndex) : fileName;
};


onMounted(() => fetchCachePreviews());
</script>

<template>
  <div :class="['uploaded-files-box', { dark: isDarkMode }]">
    <!-- <button @click="fetchcachePreviews" class="refresh-button"></button> -->
    <div v-if="Object.keys(cachePreviews).length > 0" class="file-list">
      <div v-for="(metadata, name) in cachePreviews" :key="name" class="file-item">
        <button type="button" class="file-info-button"
          style="background: none; border: none; padding: 0; margin: 0; cursor: pointer;"
          @click="handleFileClick(metadata)">
          <div class="file-info">
            <div class="file-icon">
              <font-awesome-icon :icon="faFile" />
            </div>
            <div class="button-name">
              <p class="file-name">{{ removeExtension(name) }}</p>
            </div>
          </div>
        </button>
        <button type="button" @click="deleteCache(name)" class="delete-button">
          <font-awesome-icon :icon="faTrash" />
        </button>
      </div>
    </div>
    <p v-else>No uploaded files yet.</p>
    <p v-if="error" class="error-message">{{ error }}</p>
  </div>
</template>

<style scoped>
/* Box Container */
.uploaded-files-box {
  width: 100%;
  height: 65vh;
  max-width: 200px;
  padding: 20px;
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  transition: background 0.3s ease, color 0.3s ease;
  overflow-y: auto;
  scrollbar-width: thin;
  scrollbar-color: #c1c1c1 #f1f1f1;
}

.uploaded-files-box.dark {
  background: #222;
  color: #ffffff;
  box-shadow: 0 2px 5px rgba(255, 255, 255, 0.1);
  scrollbar-color: #666 #444;
}

.uploaded-files-box::-webkit-scrollbar {
  width: 10px;
  border-radius: 10px;
}

.uploaded-files-box::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 10px;
}

.uploaded-files-box::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 100px;
}

.uploaded-files-box::-webkit-scrollbar-thumb:hover {
  background: #555;
  border-radius: 10px;
}

/* Refresh Button
.refresh-button {
  background: #007bff;
  color: white;
  padding: 8px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-weight: bold;
  margin-bottom: 10px;
  display: block;
} */

.refresh-button:hover {
  background: #005dc1;
}

.uploaded-files-box.dark .refresh-button {
  background: #555;
}

.uploaded-files-box.dark .refresh-button:hover {
  background: #777;
}

/* File List */
.file-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

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

.uploaded-files-box.dark .file-item {
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

.uploaded-files-box.dark .file-icon {
  color: #66b2ff;
}

/* File Name */
.file-name {
  font-weight: 500;
  color: #333
}

.uploaded-files-box.dark .file-name {
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

.uploaded-files-box.dark .delete-button {
  color: #ff7777;
}

.uploaded-files-box.dark .delete-button:hover {
  color: #ff9999;
}

/* Error Message */
.error-message {
  color: red;
  margin-top: 10px;
}

.refresh-button {
  width: 30px;
  height: 30px;
  justify-content: right;
  border-radius: 999px;
}
</style>