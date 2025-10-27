<script setup lang="ts">
import HeroTitle from "./HeroTitle.vue";
</script>

<template>
  <section class="hero-section">
    <HeroTitle title="Simple API" />
    <p>A project for learning Back-end with RESTful API using <a href="https://go.dev/" target="_blank">Go</a></p>

    <div class="clipboard">
      <pre>http://localhost:8080/api/v1/message</pre>
      <button type="button">Copy to Clipboard</button>
    </div>
  </section>
</template>

<style scoped>
.hero-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.clipboard {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  gap: 1rem;
}
</style>
