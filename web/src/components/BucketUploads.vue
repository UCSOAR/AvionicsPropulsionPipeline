<script setup lang="ts">
import { onMounted, reactive, ref, inject } from "vue";
import { PreviewMetadata } from "../models/PreviewMetadata";
import { endpointMapping } from "../utils/constants";
import { useMetadataStore } from "../stores/metadataStore";
import FileButton from "./FileButton.vue";

// Inject and initialize state
const isDarkMode = inject("isDarkMode", ref(false));
const cachePreviews = reactive<Record<string, PreviewMetadata>>({});
const metadataStore = useMetadataStore();
const error = ref<string | null>(null);

// Initialize objects
const objects = ref<Record<string, PreviewMetadata>>({});

const fetchObjects = async () => {
  try {
    const result = await fetch(endpointMapping.getStaticFireMetadataUrl);
    if (result.ok) {
      const data = await result.json();
      console.log(data);
      objects.value = data; // Assign data to objects
      error.value = null;
    } else {
      error.value = "Failed to fetch uploaded files.";
    }
  } catch (err) {
    error.value = (err as Error).message || "An error occurred.";
  }
};

const handleFileClick = (obj: PreviewMetadata, name: string) => {
  console.log("Tis the name", name);
  // Store metadata with the first x and y column names
  metadataStore.setMetadata(obj, obj.xColumnNames[0], obj.yColumnNames[0], name);
  console.log("Metadata stored", obj);
};

const deleteCache = async (cacheName: string) => {
  if (cacheName in cachePreviews) {
    delete cachePreviews[cacheName];
  }
};

onMounted(() => {
  fetchObjects();
});
</script>

<template>
  <div :class="['uploaded-files-box', { dark: isDarkMode }]">
    <!-- Uncomment below to add a refresh button -->
    <!-- <button @click="fetchObjects" class="refresh-button"></button> -->

    <div v-if="objects && Object.keys(objects).length > 0" class="file-list">
      <FileButton 
        v-for="(metadata, name) in objects" 
        :key="name"
        :metadata="metadata"
        :fileName="name"
        :isDarkMode="isDarkMode"
        @file-click="handleFileClick"
        @delete-click="deleteCache"
      />
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

/* Refresh Button (if used) */
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