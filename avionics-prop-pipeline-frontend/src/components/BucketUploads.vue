<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { PreviewMetadata } from '../models/BucketObject';
import { endpointMapping } from '../utils/constants';

const previews = reactive<Record<string, PreviewMetadata>>({});
const error = ref<string | null>(null);

const fetchObjects = async () => {
    try {
        const result = await fetch(endpointMapping.getStaticFireMetadataUrl);

        if (result.ok) {
            const data = await result.json();

            if (data) {
                Object.assign(previews, data);
                error.value = null;
            }
        } else {
            error.value = 'Failed to fetch objects.';
        }
    } catch (err) {
        error.value = err as string;
    }
};

const fetchColumns = async (name: string, xColumnNames: string[], yColumnNames: string[]) => {
    try {
        const result = await fetch(endpointMapping.getStaticFireColumnsUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name, xColumnNames, yColumnNames }),
        });

        if (result.ok) {
            const data = await result.json();

            if (data) {
                Object.assign(previews[name], data);
                error.value = null;
            }
        } else {
            error.value = 'Failed to fetch columns.';
        }
    } catch (err) {
        error.value = err as string;
    }
};

onMounted(() => fetchObjects());
</script>

<template>
    <div class="container">
        <button @click="fetchObjects">Refresh</button>
        <p v-if="error">{{ error }}</p>
        <div class="objects-container" v-if="Object.keys(previews).length > 0">
            <div v-for="(metadata, name) in previews" :key="name" class="object">
                <p>Cache Tree Name: {{ name }}</p>
                <p>Operator: {{ metadata.operator }}</p>
                <div class="x-column-container">
                    <button v-for="xColumn in metadata.xColumnNames" :key="xColumn"
                        @click="fetchColumns(name, [xColumn], [])">
                        {{ xColumn }}
                    </button>
                </div>
                <div class="y-column-container">
                    <button v-for="yColumn in metadata.yColumnNames" :key="yColumn"
                        @click="fetchColumns(name, [], [yColumn])">
                        {{ yColumn }}
                    </button>
                </div>
            </div>
        </div>
        <p v-else>No uploaded files.</p>
    </div>
</template>