# WordPress Docker Setup (kaas)

This folder contains a Docker setup for WordPress. You can add plugins, themes, or custom configuration as needed.

## Usage
- Edit the `Dockerfile` to customize your WordPress image.
- Add plugins/themes to the appropriate folders if needed.


### WordPress Theme Customization (kaas)

This folder contains a Docker-based setup for WordPress, focused on customizing the theme.

#### Intent
The goal is to personalize and extend the WordPress theme for your project. You can modify styles, templates, and add custom functionality as needed.

#### How to Customize the Theme

1. Edit the files in `wp-content/themes/twentytwentyfour/`:
	- Change styles in `style.css`.
	- Add or modify PHP logic in `functions.php`.
	- Update HTML templates in the `parts/` and `templates/` folders.
	- Add new patterns in the `patterns/` folder.
	- Place custom assets in `assets/` (images, fonts, CSS).

2. Build and run the WordPress container using the provided `Dockerfile` and `docker-compose.yml`.

3. Access your site and verify theme changes.

#### Tips
- Use the WordPress admin panel to activate and preview your customized theme.
- Keep theme changes versioned in Git for easy collaboration.
- Refer to the official WordPress Theme Developer Handbook for advanced customization: https://developer.wordpress.org/themes/
