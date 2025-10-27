<?php
// Custom WordPress content for demo
add_action('wp_footer', function() {
    echo '<div style="text-align:center;padding:30px;background:#ffe066;border:2px solid #ffb700;">';
    echo '<h2 style="color:#d35400;">🚀 Highlight: DevSecOps Custom Banner</h2>';
    echo '<p style="font-size:1.2em;color:#333;">This content is injected by the <strong>custom-content</strong> plugin. Edit <code>index.php</code> to change this banner.</p>';
    echo '</div>';
});
