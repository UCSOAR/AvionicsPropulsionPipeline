<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { BucketObject } from '../models/BucketObject';
import { endpointMapping } from '../utils/constants';

const objects = reactive<BucketObject[]>([]);
const error = ref<string | null>(null);

const fetchObjects = async () => {
    // try {
    //     const result = await fetch(endpointMapping.getBucketUploadsUrl);

    //     if (result.ok) {
    //         const data = await result.json();

    //         if (data) {
    //             objects.splice(0, objects.length, ...data);
    //             error.value = null;
    //         }
    //     } else {
    //         error.value = 'Failed to fetch objects.';
    //     }
    // } catch (err) {
    //     error.value = err as string;
    // }
};

const downloadObject = async (filename: string) => {
    // try {
    //     const endpoint = new URL(endpointMapping.downloadBucketObjectUrl);
    //     endpoint.searchParams.append('file', filename);

    //     const result = await fetch(endpoint);

    //     if (result.ok) {
    //         const blob = await result.blob();
    //         const url = window.URL.createObjectURL(blob);
    //         const a = document.createElement('a');

    //         a.href = url;
    //         a.download = filename;
    //         a.click();

    //         window.URL.revokeObjectURL(url);
    //     } else {
    //         error.value = 'Failed to download object.';
    //     }
    // } catch (err) {
    //     error.value = err as string;
    // }
}

onMounted(() => fetchObjects());
</script>

<template>
    <div class="container">
        <button @click="fetchObjects">Refresh</button>
        <p v-if="error">{{ error }}</p>
        <div class="objects-container" v-if="objects.length > 0">
            <div v-for="obj in objects" :key="obj.name" class="object">
                <p>File: {{ obj.name }}</p>
                <p>Size: {{ obj.size }}</p>
                <p>Last Modified: {{ obj.lastModified }}</p>
                <button @click="downloadObject(obj.name)">Download</button>
            </div>
        </div>
        <p v-else>No uploaded files.</p>
    </div>
</template>