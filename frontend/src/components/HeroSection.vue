<script setup lang="ts">
import { Icon } from "@iconify/vue";
import HeroTitle from "./HeroTitle.vue";
</script>

<template>
  <section class="hero-section">
    <HeroTitle title="Simple API" />
    <p>
      A project for learning Back-end with RESTful API using
      <a href="https://go.dev/" target="_blank">Go</a>
    </p>

    <div class="clipboard">
      <code id="api-url" class="api-url">
        http://localhost:8080/api/v1/message
      </code>
      <button type="button" class="clipboard-button">
        <Icon icon="lucide:clipboard" />
      </button>
    </div>
  </section>
</template>

<style scoped>
a {
  color: var(--color-primary-500);
}

p {
  color: var(--color-foreground-400);
  font-size: 1rem;
  text-align: center;
}

button {
  border: none;
  background-color: transparent;
}

.hero-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding-inline: 2.25rem;
}

.clipboard {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  gap: 1rem;
}

.clipboard-button {
  font-size: 1.5rem;
}

.api-url {
  background-color: var(--color-background-600);
  font-size: 0.80rem;
  border-radius: 0.2rem;
  padding-top: 0.75rem;
  padding-bottom: 0.75rem;
  padding-inline: 1.5rem;
  text-align: center;
}

@media (width >=48rem) {
  p {
    font-size: 1.25rem;
  }

  .clipboard-button {
    font-size: 1.5rem;
  }

  .api-url {
    font-size: 1rem;
  }
}
</style>
