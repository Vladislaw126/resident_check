# Resident Check PoC

> **Work in progress.** This repository contains an early-stage proof-of-concept and research draft exploring how a hypothetical resident-only eligibility rule for Dutch coffeeshops could be verified electronically while minimizing disclosure of personal data and examining the risks of profiling, logging, and future surveillance.

This project studies a practical and privacy-sensitive problem: age can often be checked from an identity document, but **current residency in the Netherlands** is harder to prove at the point of sale. Documents such as passports, ID cards, residence permits, and insurance-related cards do not map neatly onto the specific policy category that a resident-only rule would try to enforce. The project therefore focuses on the gap between **what a seller would need to know** and **what a buyer can reasonably and lawfully be asked to reveal**.

## Why this project exists

A resident-only rule for Amsterdam coffeeshops would create a practical verification problem. Sellers can usually inspect a document to estimate or verify a buyer's age, and they can identify the person standing in front of them, but the relevant policy condition is not identity alone. The condition would instead be something closer to **current residency in the Netherlands**, and that attribute is not reliably visible on the documents that all eligible residents carry.

This becomes especially difficult because Dutch residency is not represented in a single simple way across all groups. Dutch nationals, EU citizens residing in the Netherlands, and non-EU residents may all have different document situations. Some non-EU residents may hold Dutch residence documents that strongly suggest lawful residence, while many EU citizens living in the Netherlands do not carry a Dutch residence card at all. A document may reveal age or nationality without proving current residence.

A tempting shortcut would be to rely on the BSN, but BSN alone is not proof of current residency. It is therefore not enough to ask whether a number exists or appears valid. The deeper problem is how to verify the **right attribute** in a way that is technically secure, legally realistic, and privacy-preserving.

This project exists to explore whether an electronic verification system could bridge that gap. At the same time, it treats the problem critically: any infrastructure built to verify eligibility for a sensitive purchase may also create opportunities for over-collection, expanded logging, profiling, or state and institutional surveillance. The PoC is therefore as much about identifying risks and limits as it is about building a technical prototype.

## Project goal

The goal of this PoC is to explore whether a system can verify only the minimum necessary information for a hypothetical resident-only access rule, such as:

* whether the buyer is **18+**
* whether the buyer is **currently resident in the Netherlands**
* whether the verifier can reach that decision **without exposing more identity data than necessary**

The project also aims to compare a simple baseline design against a more privacy-preserving design and to analyze where security, legality, and societal concerns begin to clash.

## Main research questions

* How can current residency in the Netherlands be verified electronically without relying on documents that many eligible residents do not carry?
* Why is BSN alone insufficient as proof of residency?
* Can a verifier learn only the attributes needed for the decision, instead of learning the buyer's full identity?
* How can a system avoid storing unnecessary buyer data at the coffeeshop?
* What protections are needed against replay attacks, man-in-the-middle attacks, rogue verifier devices, and unauthorized access to sensitive data?
* How do different designs change the balance between privacy, security, and operational simplicity?
* Where could central logging, analytics, or profiling expand beyond the original purpose of the system?

## Scope

This repository currently focuses on:

* problem framing and research motivation
* early architecture exploration
* threat-model foundations
* a comparison between a **baseline central lookup design** and a **privacy-preserving attribute-verification design**
* planning a lab-based proof of concept using local hardware, virtual machines, and isolated services
* documenting security, privacy, and societal tradeoffs

## Non-goals

This project does **not** currently aim to:

* provide legal advice or claim that the proposed system would be lawful to deploy in the Netherlands
* connect to real Dutch government registries or production identity systems
* build a production-ready verification platform
* define actual public policy for coffeeshops or cannabis regulation
* prove that resident-only enforcement is desirable as a matter of policy
* collect or process real personal data

## Why this is difficult

The problem is difficult because the policy category that might need to be enforced does not align neatly with the proofs people carry in daily life. Identity, age, lawful stay, and current residency are related, but they are not the same thing.

It is also difficult because the most obvious technical shortcut, a central identifier check, is not automatically the right design. Even if a number is valid, that does not mean it proves the attribute that matters. In addition, any architecture that centralizes sensitive verification requests creates a larger attack surface, more logging opportunities, and stronger incentives for future misuse.

Finally, this is difficult because the project sits at the intersection of security engineering, privacy engineering, and administrative reality. A technically simple system may be socially dangerous. A privacy-preserving system may be harder to deploy or explain. A secure system may still normalize sensitive forms of monitoring.

## Privacy and societal concerns

This project treats privacy and societal impact as core design concerns rather than side notes. A verification system for a sensitive purchase category could gradually expand beyond its original purpose. Even if initially designed only to answer an eligibility question, the same infrastructure might later be used for richer logging, behavioral analysis, profiling, or public-health monitoring that identifies groups or individuals more than originally intended.

For that reason, the PoC will pay special attention to data minimisation, limited disclosure, auditability, and the separation between what is necessary for an eligibility decision and what would be convenient for analytics or enforcement. One of the central questions of the project is whether technical design can meaningfully resist surveillance creep rather than merely making it more efficient.

## Current approach

The current working approach is to compare two design directions.

The first is a **baseline architecture**, in which a verifier terminal sends a buyer-linked identifier or request to a backend that returns some form of allow/deny decision. This design is useful because it is simple to reason about and easy to prototype, but it raises strong privacy and misuse concerns.

The second is a **privacy-preserving architecture**, in which the verifier learns only the attributes required for the decision, such as `18+` and `resident in NL`, rather than receiving the buyer's full identity or a raw identifier. This design is more aligned with selective disclosure and minimal-data principles, but it is more complex from a protocol and system-design perspective.

The PoC will likely use both as comparison points.

## Repository structure

The exact layout is still evolving, but the repository will likely include the following main areas:

* `docs/` — research draft, problem statement, architecture notes, threat-model notes, and report planning
* `diagrams/` — system context, container, trust-boundary, and sequence diagrams
* `apps/` — verifier, backend, issuer, wallet simulator, or other PoC components
* `services/` — mock registry, logging, key-management, or revocation-related services
* `tests/` — integration, end-to-end, security, and performance tests
* `infra/` — VM, container, device, and local lab setup notes
* `telemetry/` — optional observability and privacy-filter experiments
* `report/` — later-stage material for the final article, report, or figures

## Status

Current status: **Milestone 0 — research framing and problem specification**.

At this stage, the repository is being used to clarify the motivation, narrow the problem, distinguish between similar-but-different concepts such as identity and residency, and prepare the backbone for later architecture work, threat modeling, and implementation.

## Planned milestones

1. Research framing and problem specification
2. Threat-model foundation and trust-boundary identification
3. Architecture drafting and protocol comparison
4. Lab setup and hardware planning
5. Baseline PoC implementation
6. Security testing of the baseline design
7. Privacy-preserving PoC implementation
8. Comparative analysis and benchmarking
9. Optional telemetry and observability experiments
10. Final report/article preparation

## Legal and ethical note

This repository is a **research proof of concept**. It is not a production-ready system and should not be interpreted as a claim that such a system would be legally deployable, socially desirable, or ethically justified in its current form.

The project explicitly studies the tension between enforceability, privacy, and surveillance risk. Any later implementation or publication should preserve that critical perspective and avoid normalizing unnecessary collection of sensitive personal data.

## Further documentation

Planned deeper documentation will live under `docs/`, including files such as:

* `notes/00_research_draft.md`
* `report/sections/01_problem_statement.md`
* `report/sections/02_scope_and_nongoals.md`
* `report/sections/03_threat_model_notes.md`
* `report/sections/04_architecture_notes.md`
* `report/sections/05_report_summary.md`

These documents will gradually evolve into the foundation for the later report.
