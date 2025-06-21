export const USER_TYPES = {
  CLIENT: 0,
  AGENT: 1,
  SALES_CHANNEL: 2,
  INFLUENCER: 3,
  DISTRIBUTOR: 4,
  DOCTOR: 5,
  OPERATOR: 10,
  ADMIN: 11,
  SUPER_ADMIN: 12,
} as const;

export const USER_TYPE_LABELS = {
  [USER_TYPES.CLIENT]: 'Patient',
  [USER_TYPES.AGENT]: 'Agent',
  [USER_TYPES.SALES_CHANNEL]: 'Sales Channel',
  [USER_TYPES.INFLUENCER]: 'Influencer',
  [USER_TYPES.DISTRIBUTOR]: 'Distributor',
  [USER_TYPES.DOCTOR]: 'Doctor',
  [USER_TYPES.OPERATOR]: 'Operator',
  [USER_TYPES.ADMIN]: 'Admin',
  [USER_TYPES.SUPER_ADMIN]: 'Super Admin',
} as const;

export const MEDICAL_SPECIALTIES = [
  {
    id: 1,
    title: "Cancer Immunotherapy",
    description: "CAR-T Cell Therapy, BiTE Antibodies, and Neoantigen treatments",
    icon: "🔬",
    treatments: ["CAR-T Cell Therapy", "BiTE Antibody", "Neoantigen-TIL"]
  },
  {
    id: 2,
    title: "Autoimmune Disorders",
    description: "Precision therapy for Psoriasis, Rheumatoid Arthritis, and Lupus",
    icon: "🛡️",
    treatments: ["Psoriasis Treatment", "RA Therapy", "Lupus Management"]
  },
  {
    id: 3,
    title: "Ophthalmology",
    description: "Advanced treatments for eye diseases and vision disorders",
    icon: "👁️",
    treatments: ["Glaucoma Treatment", "Macular Degeneration", "Optic Nerve Therapy"]
  },
  {
    id: 4,
    title: "Neurological Sciences",
    description: "Stroke rehabilitation, Alzheimer's, and neuroplasticity",
    icon: "🧠",
    treatments: ["Stroke Rehab", "Alzheimer's Care", "Autism Intervention"]
  },
  {
    id: 5,
    title: "Respiratory Medicine",
    description: "Tuberculosis and pneumonia precision treatments",
    icon: "🫁",
    treatments: ["TB Treatment", "Pneumonia Therapy", "Bronchitis Care"]
  },
  {
    id: 6,
    title: "Infectious Diseases",
    description: "HPV immunotherapy and antibiotic-resistant infections",
    icon: "🦠",
    treatments: ["HPV Immunotherapy", "Resistant Infections", "Antiviral Protocols"]
  }
];
