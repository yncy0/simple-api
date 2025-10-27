<script setup lang="ts">
interface Props {
  title: string
}

const props = defineProps<Props>();
</script>

<template>
  <h1>
    {{ props.title }}
  </h1>
</template>
