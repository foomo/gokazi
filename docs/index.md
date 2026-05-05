---
layout: home

hero:
  name: gokazi
  text: Daemonless process manager
  tagline: 
  image:
    src: /logo.png
    alt: gokazi
  actions:
    - theme: brand
      text: Get started
      link: /guide/introduction
    - theme: alt
      text: View on GitHub
      link: https://github.com/foomo/gokazi

features:
  - title: Daemonless
    details: No supervisor, no PID file. Every command enumerates OS processes and matches by name, path, cwd, and args.
  - title: YAML-driven
    details: Declare your tasks once. Layer environment-specific overrides via repeated <code>-c</code> flags.
  - title: Cross-platform
    details: macOS, Linux, and Windows binaries. Install via Homebrew, Docker, mise, or <code>go install</code>.
  - title: Make / Just friendly
    details: gokazi is the stop button, not the start button. Drop it into your existing dev scripts.
---
