'use client';

import { Metadata } from "next";
import React, { useEffect, useRef } from "react";
import styles from './scroll.module.css';

export const metadata: Metadata = {
  title: "Template Manager",
  description: "Manage your templates with ease",
};


export default function Scroll() {
  const animatedTextRef = useRef<HTMLDivElement>(null);


  useEffect(() => {
    // Start animation when component mounts
    if (animatedTextRef.current) {
      const textContainer = animatedTextRef.current;
      const textWidth = textContainer.offsetWidth;
      const containerWidth = textContainer.parentElement?.offsetWidth || 0;
      const duration = (textWidth + containerWidth) * 10; // Adjust speed based on width

      textContainer.style.animationDuration = `${duration}ms`;
    }
  }, []);

  return (
      <div className={styles.scrollableText} ref={animatedTextRef}>
        Why Choose Us? {/* Text to animate */}
      </div>
  );
}
