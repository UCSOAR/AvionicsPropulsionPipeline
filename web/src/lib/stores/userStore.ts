import { writable } from "svelte/store";
import type { User } from "$lib/models/user";

export const userStore = writable<User | null>(null);
