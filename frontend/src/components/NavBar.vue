<template>
    <div>
        <nav>
            <ul>
                <li><RouterLink to="/">Home</RouterLink></li>
                <li><RouterLink to="/about">About</RouterLink></li>
            </ul>
        </nav>
        <slot></slot>
    </div>
</template>

<script>
    import { RouterLink } from "vue-router";
</script>