<?php
// Make WordPress see HTTPS when behind ingress
if (!empty($_SERVER['HTTP_X_FORWARDED_PROTO'])
    && stripos($_SERVER['HTTP_X_FORWARDED_PROTO'], 'https') !== false) {
    $_SERVER['HTTPS'] = 'on';
}

// Force canonical public URLs (before DB options are read)
add_filter('pre_option_home', function () { return 'https://devsecops2025-arubacloud.com'; });
add_filter('pre_option_siteurl', function () { return 'https://devsecops2025-arubacloud.com'; });

// Ensure admin uses HTTPS
if (!defined('FORCE_SSL_ADMIN')) {
    define('FORCE_SSL_ADMIN', true);
}
