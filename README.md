# Pennypecker

> **"A Terraform provider that pecks away at cloud costs by blocking deployments that exceed budgets or trigger anomalies."**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Terraform](https://img.shields.io/badge/Terraform-1.0+-7B42BC?style=flat&logo=terraform)](https://www.terraform.io/)
[![License](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)

---

## 🚧 Work in Progress

This provider is currently under active development. Check back soon for updates!

---

## 🎯 Vision

Enable **proactive cost governance** by exposing live AWS cost data, forecasts, and anomaly detection directly into your Terraform workflows — blocking costly deployments before they happen.

---

## 📦 Planned Features

- **Data Sources**
  - `pennypecker_current_spend` - Query actual AWS spend
  - `pennypecker_forecast` - Get cost forecasts
  - `pennypecker_anomalies` - Detect cost anomalies
  - `pennypecker_budget_remaining` - Check budget headroom
  - `pennypecker_ri_coverage` - Reserved Instance coverage
  - `pennypecker_savings_plan_coverage` - Savings Plan coverage

- **Resources**
  - `pennypecker_cost_policy` - Define cost enforcement rules
  - `pennypecker_savings_plan` - Purchase Savings Plans

---

## 🔧 Building from Source

```bash
git clone https://github.com/avinash-exp/pennypecker.git
cd pennypecker
make build
```

---

## 📄 License

This project is licensed under the Mozilla Public License 2.0 - see the [LICENSE](LICENSE) file for details.