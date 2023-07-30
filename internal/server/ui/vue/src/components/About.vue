<template>
  <div class="mx-8 my-8" style="white-space: pre-wrap;">
    <h1 class="text-h5"><b>Обо мне</b></h1>
    <p class="my-3"> Привет, меня зовут Серега (можно Сегей или Саргей, так я случайно писал свое имя в билетах на
      самолет).
      Мой путь начался с разработки драйверов под микроконтроллеры STM32 на языке С,
      а сейчас я сеньор с 5 годами опыта на Goland, NodeJs и Typescript который обитает на просторах
      Чудесного Бали
    </p>
    <h1 class="text-h5"><b>Актуальность и экспертиза</b></h1>
    <div class="my-3">
      Проект развивается и дополняется актуальным функционалом благодаря Вадиму, человеком - имеющим
      успешный опыт в охоте за airdrop (
      <a href="https://community.optimism.io/docs/governance/airdrop-1/" target="_blank">Optimism</a>,
      <a href="https://docs.arbitrum.foundation/airdrop-eligibility-distribution" target="_blank">Arbitrum</a>,
      SUI
      ) c опытом порядка 2 лет. Мы работаем в партнерстве, Вадим задает курс - я занимаюсь автоматизацией
    </div>
    <h1 class="text-h5"><b>Цель</b></h1>
    <div class="my-3">
      В сообществах по <a
      href="https://en.wikipedia.org/wiki/Airdrop_(cryptocurrency)" target="_blank">airdrop hunting</a>
      множество скриптов, всяких разных: консольных, на python, на nodejs, с интерфейсом в Telegram и
      Discord, даже десктопные версии есть.
      Это здорово, но на мой профессиональный и пользовательский взгляд не вполне удобно.
      Особенно когда большая часть пользователей не являются профессиональными программистами.
      Поэтому я решил сделать продукт, который был бы удобен в использовании и не заставлял пользователей
      тратить время на изучение языков программирования.
    </div>
    <h1 class="text-h5"><b>Скам</b></h1>
    <div class="my-3">
      Риски всегда есть, везде и во всем. Забавно, что я не видел на скам сайтах разделов про скам. Итак о тарелочках:
      <h1 class="text-h6"><b>Человеческие риски</b></h1>
      <p>
        Проект разрабатывается, поддерживается и амнистируется одним человеком (мной). Мне и Вадиму проект позволил
        маштабироваться и упростил
        жизнь в погоне за дропами. В ходе разработки я понял что получается удобный интерфейс. Планирую дальше развивать
        и масштабировать свое детище.
        моя страсть и интерес - разработка ПО, это мое хобби и работа.
        На мой взгляд это уменьшает вероятность скама по человеческому фактору.
        <span class="text-disabled">PS. разработчик с моим опытом, стажем и стеком в среднем получает 5k$ в месяц.
          Живу я один, живу скромно, деньги с ваших кошелей меня не особо интересуют.</span>
      </p>
      <h1 class="text-h6"><b>Риски взлома</b></h1>
      <p>
        Для разработки приложения я использовал актуальные современные практики и технологии (
        <a href="https://go.dev/" target="_blank">Golang</a>,
        <a href="https://nodejs.org/en" target="_blank">NodeJs</a>,
        <a href="https://vuejs.org/" target="_blank">VueJs 3</a>,
        <a href="https://www.typescriptlang.org/" target="_blank">TypeScript</a>,
        <a href="https://www.postgresql.org/" target="_blank">Postgres 15.1</a>, GRPC ).
        Все чувствительные данные шифруются
        с помощью <a href="https://en.wikipedia.org/wiki/Advanced_Encryption_Standard" target="_blank">AES</a> алгоритма
        шифрования. БД закрыта белым списком ip + ssl,
        Даже если бд угонят - ничего из нее не получат.

        Пример: приватный ключ <b>eee5119cd7243f855896004e9cb9c8c1eeb3d9f8dac388d4901dbbf3dec683aa</b>
        хранится в базе данных в виде массива байтов:
        <b>0x3E6F210F5AC501BD71BD984F8B1C7A3301BFAAF36A8E5CF39ECDA9AA92B2D47BC324D80A059CF4A06EF4D89871E
          67BFCFE47FE8C5A25FDAE4E95BFEF14535B9EE6D619779631220E5518EED6180CE07E9D5A04CB15939E5D4D7A7E5E</b>.
        Для взлома может уйти не один век.
        Все логины пользователей изолированы друг от друга, никто, даже я - не может посмотреть сколько денег у других
        пользователей.
        Супер-админ-аккаунта нет и не будет.


        Хостинг сервера осуществляется на <a href="https://www.digitalocean.com/" target="_blank">DigitalOcean</a> - это
        один из наиболее известных cloud провайдеров для хостинга серверов с отличным аптаймом. Работаю с ним не один
        год.
        Проект контейнеризирован с помощью <a href="https://www.docker.com/" target="_blank">Docker</a> и развернут в
        локальной сети на сервере с помощью
        <a href="https://docs.docker.com/compose/" target="_blank">Docker-compose</a>. Это гарантирует изолированность
        приложения от <a href="https://en.wikipedia.org/wiki/Virtual_private_server" target="_blank">VPS</a>.
        Все сетевые порты, кроме стандартного порта для web приложений и ssh туннеля закрыты. Доступ на сервер
        возможен только с использованием
        <a href="https://en.wikipedia.org/wiki/Secure_Shell" target="_blank">ssh</a> ключа, который есть только у меня.
        Аутентификация в приложении возможна только через <a
        href="https://developers.google.com/identity/protocols/oauth2?hl=ru" target="_blank">Oauth</a> протокол через
        провайдера Google. Для безопасности введена сущности Profile и Account - благодаря ним гарантируется, что не
        будут логироваться приложением.
        Ксли логи сольют - ничего важного из них не узнают.
        чувствительные данные. Cервер закрыт nginx с принудительным оборачиваением трафика через SSl. Общение
        браузер-сервер дополнительно
        шифруется сначала в AES потом в BASE64.

      </p>
    </div>
    <a href="https://t.me/shinybozdo" class="text-blue" target="_blank">мой Telegram</a>
  </div>

</template>

<script lang="ts">

import {defineComponent} from 'vue';

export default defineComponent({
  name: "About",
})


</script>


<style>
</style>

