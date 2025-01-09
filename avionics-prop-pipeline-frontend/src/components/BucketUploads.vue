<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { BucketObject } from '../models/BucketObject';
import { endpointMapping } from '../utils/constants';

const objects = reactive<BucketObject[]>([]);
const fetchObjectError = ref<string | null>(null);

const fetchObjects = async () => {
    try {
        const result = await fetch(endpointMapping.getBucketUploadsUrl);

        if (result.ok) {
            const data = await result.json();

            objects.splice(0, objects.length, ...data);
            fetchObjectError.value = null;
        } else {
            fetchObjectError.value = 'Failed to fetch objects.';
        }
    } catch (err) {
        fetchObjectError.value = err as string;
    }
};

onMounted(() => fetchObjects());
</script>

<template>
    <div class="container">
        <button @click="fetchObjects">Refresh</button>
        <p v-if="fetchObjectError">{{ fetchObjectError }}</p>
        <div class="objects-container" v-if="objects.length > 0">
            <div v-for="obj in objects" :key="obj.name" class="object">
                <p>File: {{ obj.name }}</p>
                <p>Size: {{ obj.size }}</p>
                <p>Last Modified: {{ obj.lastModified }}</p>
            </div>
        </div>
        <p v-else>No uploaded files.</p>
    </div>
</template>