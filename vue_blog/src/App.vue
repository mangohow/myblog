<template>
    <router-view/>
</template>

<style>

body {
    width: 100%;
    height: 100%;
    overflow-x: hidden;
}

::-webkit-scrollbar{
    width: 6px !important;
    height: 5px !important;
    border-radius: 10px !important;
    background-color: #F5F5F5 !important;
}


::-webkit-scrollbar-thumb {
    background-image: linear-gradient(to top, #fa709a 0%, #fee140 100%) !important;
    background-clip:padding-box !important;
    -webkit-border-radius: 10em !important;
    -moz-border-radius: 10em !important;
    border-radius: 10em !important;
    transition: 0.3s ease-in-out !important;
}
::-webkit-scrollbar-track-piece {
    background-color:transparent !important;
    -webkit-border-radius: 0 !important;
    -moz-border-radius: 0 !important;
    border-radius: 0 !important;
}
::-webkit-scrollbar-thumb:hover {
    background-image: linear-gradient(120deg, #f093fb 0%, #f5576c 100%) !important;
    transition: 0.3s ease-in-out !important;
}


</style>
