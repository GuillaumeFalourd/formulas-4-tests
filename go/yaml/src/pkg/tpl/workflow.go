package tpl

const (
	Workflow = `
name: Generate Datas

on:
  workflow_dispatch: # execute manually
  schedule: # execute every 12 hours
    - cron: "* */12 * * *"

jobs:

  summary-cards:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - name: Summary Cards
        uses: vn7n24fzkq/github-profile-summary-cards@release
        env:
          GITHUB_TOKEN: {{.GITHUB_TOKEN}}
        with:
          USERNAME: {{.REPO_OWNER}}
          
  snake-animation:
    runs-on: ubuntu-latest
    steps:
      - name: Generate Snake Animation
        uses: Platane/snk@master
        id: snake-gif
        with:
          github_user_name: {{.username}}
          svg_out_path: dist/github-contribution-grid-snake.svg
          
      - name: Publish Snake Animation SVG
        uses: crazy-max/ghaction-github-pages@v2.1.3
        with:
          target_branch: output
          build_dir: dist
        env:
          GITHUB_TOKEN: {{.GITHUB_TOKEN}}

  metrics-achievements:
    runs-on: ubuntu-latest
    steps:
      - name: Achievements metrics
        uses: lowlighter/metrics@master
        with:
          filename: metrics.plugin.achievements.svg
          token: {{.ACCESS_TOKEN}}
          base: ""
          plugin_achievements: yes
          plugin_achievements_threshold: B       # Display achievements with rank B or higher
          plugin_achievements_secrets: yes       # Display unlocked secrets achievements
          plugin_achievements_ignored: octonaut  # Hide octonaut achievement
          plugin_achievements_limit: 0           # Display all unlocked achievement matching threshold and secrets params

  metrics-github:
    runs-on: ubuntu-latest
    steps:
      - name: Github metrics
        uses: lowlighter/metrics@master
        with:
          token: {{.ACCESS_TOKEN}}

`
)
