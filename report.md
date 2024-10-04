#### Tableau de Synthèse des Tickets

| N° Ticket | Client       | Sujet                                                                                                    | Statut | Équipe               | Heure de Détection | Criticité | Pourquoi Urgent                                      |
|-----------|--------------|----------------------------------------------------------------------------------------------------------|--------|----------------------|--------------------|-----------|----------------------------------------------------|
| 1661617   | AVRIL        | AVR-SRV-FRDCASNA-0032 Équipement Critique Oracle Failed Jobs - Base FRDCASNA-0032_KLB00T1                | 1      | (Pôle) IFG-SQUAD-1   | 23:55              | 1         | Oracle toujours critique                            |
| 1661614   | AVRIL        | AVR-SRV-FRDCASPDBS-016 Équipement Critique CHECK-ALL-UNIX-DISK                                           | 1      | (Pôle) IFG-SQUAD-1   | 23:08              | 2         | Dépassement capacité disque                         |
| 1661612   | NMMTN - Valoty | VAL-UVE-AZA / UPTIME                                                                                    | 1      | (Pôle) IFG-SQUAD-2-3-4 | 21:52              | 2         | Délais de résolution comprimé                      |
| 1661611   | IRD          | IRD-SRV-CENSVRAP9243 / UNIX-TIME                                                                         | 1      | (Pôle) IFG-SQUAD-2-3-4 | 21:30              | 2         | Délais de résolution comprimé                      |
| 1661608   | IRD          | IRD-SRV-CENSVRAP9243 / SSH                                                                               | 1      | (Pôle) IFG-SQUAD-2-3-4 | 21:12              | 2         | Délais de résolution comprimé                      |
| 1661594   | AVRIL        | Autre demande - mettre un boitier avec un switch administrable et une baie de brassage                   | 1      | (Pôle) IFG-SQUAD-1   | 17:52              | 3         | Demande hors catalogue                              |
| 1661593   | SPHERE       | SPHERE-SRV-AD02 / WINDOWS-SNMP-CPU                                                                       | 1      | (Pôle) IFG-SQUAD-2-3-4 | 17:50              | 5         | SNMP moins critique                                 |
| 1661588   | AVRIL        | Migration DS SNC - Site PAR7s - En HNO                                                                    | 1      | (Pôle) IFG-SQUAD-1_BACKOFFICE | 17:15      | 2         | Dépassement délai intervention                      |
| 1661586   | CNG          | [EDN] Log websocket                                                                                       | 1      | (Pôle) IFG-SQUAD-2-3-4 | 17:12              | 3         | Demande service                                     |
| 1661585   | AVRIL        | Tâche transverse                                                                                         | 1      | (Pôle) IFG-SQUAD-1   | 17:11              | 3         | Demande interne                                     |

#### Plan d'Action pour les Criticités 1

1. **Ticket N° 1661617 - AVRIL**
   - **Sujet** : Équipement Critique Oracle Failed Jobs - Base FRDCASNA-0032_KLB00T1
   - **Action** : Mobiliser immédiatement l'équipe DBA pour intervention prioritaire sur Oracle.

2. **Ticket N° 1661614 - AVRIL**
   - **Sujet** : Équipement Critique CHECK-ALL-UNIX-DISK
   - **Action** : Vérifier et gérer l'utilisation du disque pour éviter des défaillances critiques. Planifier une intervention rapide.

Ces actions permettront de focaliser les ressources sur les problèmes les plus impactants pour la continuité des services critiques. Cela inclut des interventions immédiates sur les systèmes Oracle et la gestion des ressources disque pour prévenir les pannes.

#### Critères d'Urgence et Criticité Définis
- **Oracle** : Toujours critique en raison de son importance stratégique.
- **SNMP** : Moins critique, car souvent utilisé pour de la surveillance non essentielle.
- **Dépassement Capacité Disque** : Élevée car peut causer des interruptions de service.
- **Délais de Résolution** : Important pour respecter les engagements de service (SLA).
- **Demande hors catalogue** : Moins critique, sauf si impact opérationnel direct.

Les tickets sont classés et les plans d'action sont proposés pour assurer une gestion efficace et prioritaire des incidents les plus critiques.